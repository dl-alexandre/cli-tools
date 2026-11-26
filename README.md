# cli-tools

[![CI](https://github.com/dl-alexandre/cli-tools/actions/workflows/ci.yml/badge.svg)](https://github.com/dl-alexandre/cli-tools/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dl-alexandre/cli-tools)](https://goreportcard.com/report/github.com/dl-alexandre/cli-tools)
[![Go Reference](https://pkg.go.dev/badge/github.com/dl-alexandre/cli-tools.svg)](https://pkg.go.dev/github.com/dl-alexandre/cli-tools)

`cli-tools` contains shared, focused helpers for Go CLI applications. It targets Go 1.26 and keeps its exported APIs small so each CLI can retain its own command framework and presentation layer.

## Packages

- **`version`** — build metadata, version formatting, normalization, and comparison.
- **`update`** — GitHub latest-release checks, cached results, and update notices.
- **`cache`** — JSON file cache and thread-safe in-memory LRU cache.
- **`retry`** — context-aware operation retries with exponential backoff, jitter, and HTTP `Retry-After` support.
- **`appleauth`** — ES256 JWT signing from PEM or key files, with thread-safe token caching and early refresh.

## Install

```sh
go get github.com/dl-alexandre/cli-tools@v0.0.1
```

## Quick start

```go
package main

import (
	"fmt"

	"github.com/dl-alexandre/cli-tools/version"
)

func main() {
	fmt.Println(version.String())
}
```

Set build metadata with linker flags:

```sh
go build -ldflags "-X github.com/dl-alexandre/cli-tools/version.Version=v1.2.3 -X github.com/dl-alexandre/cli-tools/version.BinaryName=myapp"
```

## Examples

### Retry a transient operation

```go
result := retry.Do(ctx, retry.Config{
	MaxAttempts: 4, // includes the first attempt
}, func(ctx context.Context) retry.Result[*http.Response] {
	resp, err := client.Do(request.WithContext(ctx))
	if err != nil {
		return retry.Result[*http.Response]{Err: err, Retry: isTemporary(err)}
	}
	if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500 {
		wait, _ := retry.RetryAfter(resp.Header.Get("Retry-After"))
		return retry.Result[*http.Response]{Value: resp, Retry: true, RetryAfter: wait}
	}
	return retry.Result[*http.Response]{Value: resp}
})
if result.Err != nil {
	return result.Err
}
```

`Retry-After` accepts delay seconds and HTTP dates. The operation decides which errors and response codes are safe to retry.

### Sign an Apple ES256 token

```go
signer, err := appleauth.NewFromFile(appleauth.Config{
	KeyID:    "ABC123DEFG",
	TeamID:   "DEF123GHIJ", // used as iss when Issuer is not set
	Audience: "appstoreconnect-v1",
	TTL:      15 * time.Minute,
}, "AuthKey_ABC123DEFG.p8")
if err != nil {
	return err
}

token, err := signer.Token() // reused until the early-refresh window
```

`appleauth.New(config, pemBytes)` accepts key bytes directly. `Issuer` takes precedence over `TeamID` for the JWT `iss` claim.

### Cache a value

```go
c := cache.New(cache.DefaultDir("myapp"), 24*time.Hour)
if err := c.Set("response", payload, time.Hour); err != nil {
	return err
}
```

### Check for updates

```go
checker := update.New(update.Config{
	CurrentVersion: version.Version,
	BinaryName:     "myapp",
	GitHubRepo:     "owner/myapp",
})
info, err := checker.Check(false)
if err != nil {
	return err
}
return update.DisplayUpdate(info, "myapp", "table", "brew upgrade myapp")
```

## Development

```sh
go mod tidy
go test -race ./...
go vet ./...
golangci-lint run
```

For compatibility, the library keeps the existing `version.Version`, `version.String()`, `update.New(...)`, and cache APIs stable.
