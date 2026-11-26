package retry

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestDoRetriesAndReturnsSuccess(t *testing.T) {
	var calls atomic.Int32
	got := Do(context.Background(), Config{
		MaxAttempts: 3, InitialDelay: time.Millisecond, DisableJitter: true,
	}, func(context.Context) Result[string] {
		if calls.Add(1) < 3 {
			return Result[string]{Err: errors.New("temporary"), Retry: true}
		}
		return Result[string]{Value: "ok"}
	})
	if got.Err != nil || got.Value != "ok" || calls.Load() != 3 {
		t.Fatalf("Do() = %#v after %d calls", got, calls.Load())
	}
}

func TestDoStopsWhenOperationDeclinesRetry(t *testing.T) {
	var calls int
	wantErr := errors.New("permanent")
	got := Do(context.Background(), Config{MaxAttempts: 4}, func(context.Context) Result[int] {
		calls++
		return Result[int]{Err: wantErr}
	})
	if !errors.Is(got.Err, wantErr) || calls != 1 {
		t.Fatalf("Do() = %#v after %d calls", got, calls)
	}
}

func TestDoHonorsRetryAfter(t *testing.T) {
	var calls int
	got := Do(context.Background(), Config{MaxAttempts: 2}, func(context.Context) Result[string] {
		calls++
		if calls == 1 {
			return Result[string]{Retry: true, RetryAfter: time.Millisecond}
		}
		return Result[string]{Value: "retried"}
	})
	if got.Value != "retried" || calls != 2 {
		t.Fatalf("Do() = %#v after %d calls", got, calls)
	}
}

func TestDoContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var calls atomic.Int32
	go func() {
		for calls.Load() == 0 {
			time.Sleep(time.Millisecond)
		}
		cancel()
	}()
	got := Do(ctx, Config{MaxAttempts: 5}, func(context.Context) Result[struct{}] {
		calls.Add(1)
		return Result[struct{}]{Retry: true, RetryAfter: time.Hour}
	})
	if !errors.Is(got.Err, context.Canceled) || calls.Load() != 1 {
		t.Fatalf("Do() = %#v after %d calls", got, calls.Load())
	}
}

func TestRetryAfterParsesSecondsAndHTTPDate(t *testing.T) {
	now := time.Date(2026, 11, 26, 12, 0, 0, 0, time.UTC)
	for _, tt := range []struct {
		name   string
		header string
		want   time.Duration
		ok     bool
	}{
		{name: "seconds", header: " 12 ", want: 12 * time.Second, ok: true},
		{name: "date", header: now.Add(30 * time.Second).Format(http.TimeFormat), want: 30 * time.Second, ok: true},
		{name: "past date", header: now.Add(-time.Minute).Format(http.TimeFormat), want: 0, ok: true},
		{name: "invalid", header: "later", ok: false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := RetryAfterAt(tt.header, now)
			if got != tt.want || ok != tt.ok {
				t.Fatalf("RetryAfterAt() = (%v, %v), want (%v, %v)", got, ok, tt.want, tt.ok)
			}
		})
	}
}

func TestRetryAfterHeaderFromHTTPResponse(t *testing.T) {
	response := httptest.NewRecorder()
	response.Header().Set("Retry-After", "3")
	got, ok := RetryAfter(response.Header().Get("Retry-After"))
	if !ok || got != 3*time.Second {
		t.Fatalf("RetryAfter() = (%v, %v)", got, ok)
	}
}
