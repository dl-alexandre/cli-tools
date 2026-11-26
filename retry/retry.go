// Package retry runs operations again after transient failures.
package retry

import (
	"context"
	"errors"
	"math/rand/v2"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ErrNilOperation is returned when Do receives no operation.
var ErrNilOperation = errors.New("retry: operation is nil")

// Config controls retry timing. MaxAttempts includes the first call. Zero or
// negative values use the defaults documented below.
type Config struct {
	// MaxAttempts includes the first call; the default is 3.
	MaxAttempts int
	// InitialDelay is the delay before the first retry; the default is 100ms.
	InitialDelay time.Duration
	// MaxDelay bounds exponential delays; the default is 5s.
	MaxDelay time.Duration
	// Multiplier scales the delay after each attempt; the default is 2.
	Multiplier float64
	// Jitter is the random fraction applied to both sides of a delay; the
	// default is 0.2. Set DisableJitter to true for fixed delays.
	Jitter float64
	// DisableJitter turns off randomizing the backoff delay.
	DisableJitter bool
}

// Result is the outcome of one operation attempt. Set Retry when Err is
// transient, and set RetryAfter to honor a server-provided delay.
type Result[T any] struct {
	Value      T
	Err        error
	Retry      bool
	RetryAfter time.Duration
}

// Do invokes operation until it succeeds, declines a retry, the context is
// canceled, or MaxAttempts is reached. Defaults are 3 attempts, 100ms initial
// delay, 5s maximum delay, a 2x multiplier, and up to 20% jitter.
func Do[T any](ctx context.Context, cfg Config, operation func(context.Context) Result[T]) Result[T] {
	if ctx == nil {
		ctx = context.Background()
	}
	if operation == nil {
		var zero T
		return Result[T]{Value: zero, Err: ErrNilOperation}
	}
	cfg = defaults(cfg)
	delay := cfg.InitialDelay
	var last Result[T]

	for attempt := 0; attempt < cfg.MaxAttempts; attempt++ {
		if err := ctx.Err(); err != nil {
			return Result[T]{Err: err}
		}
		last = operation(ctx)
		if !last.Retry || attempt == cfg.MaxAttempts-1 {
			return last
		}

		wait := last.RetryAfter
		if wait <= 0 {
			wait = jitter(delay, cfg.Jitter)
			if wait > cfg.MaxDelay {
				wait = cfg.MaxDelay
			}
		}
		if err := sleep(ctx, wait); err != nil {
			return Result[T]{Err: err}
		}
		delay = time.Duration(float64(delay) * cfg.Multiplier)
		if delay > cfg.MaxDelay || delay < 0 {
			delay = cfg.MaxDelay
		}
	}
	return last
}

// RetryAfter parses the HTTP Retry-After header as seconds or an HTTP date.
// The returned duration is zero and ok is false for malformed values.
func RetryAfter(header string) (duration time.Duration, ok bool) {
	return RetryAfterAt(header, time.Now())
}

// RetryAfterAt is RetryAfter with an explicit clock, useful for deterministic
// callers and tests.
func RetryAfterAt(header string, now time.Time) (time.Duration, bool) {
	header = strings.TrimSpace(header)
	if header == "" {
		return 0, false
	}
	if seconds, err := strconv.ParseInt(header, 10, 64); err == nil {
		if seconds < 0 || seconds > int64((time.Duration(1<<63-1))/time.Second) {
			return 0, false
		}
		return time.Duration(seconds) * time.Second, true
	}
	at, err := http.ParseTime(header)
	if err != nil {
		return 0, false
	}
	if !at.After(now) {
		return 0, true
	}
	return at.Sub(now), true
}

func defaults(cfg Config) Config {
	if cfg.MaxAttempts <= 0 {
		cfg.MaxAttempts = 3
	}
	if cfg.InitialDelay <= 0 {
		cfg.InitialDelay = 100 * time.Millisecond
	}
	if cfg.MaxDelay <= 0 {
		cfg.MaxDelay = 5 * time.Second
	}
	if cfg.InitialDelay > cfg.MaxDelay {
		cfg.InitialDelay = cfg.MaxDelay
	}
	if cfg.Multiplier <= 1 {
		cfg.Multiplier = 2
	}
	if cfg.Jitter < 0 {
		cfg.Jitter = 0
	} else if cfg.Jitter == 0 && !cfg.DisableJitter {
		cfg.Jitter = 0.2
	}
	if cfg.Jitter > 1 {
		cfg.Jitter = 1
	}
	return cfg
}

func jitter(delay time.Duration, fraction float64) time.Duration {
	if fraction == 0 || delay <= 0 {
		return delay
	}
	factor := 1 - fraction + rand.Float64()*2*fraction
	return time.Duration(float64(delay) * factor)
}

func sleep(ctx context.Context, delay time.Duration) error {
	if delay <= 0 {
		return nil
	}
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
