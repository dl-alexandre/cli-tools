# Migration Guide

`cli-tools` provides version metadata, update checks, caching, retry policy, and Apple ES256 signing. Add only the packages your CLI needs.

## Add the module

```sh
go get github.com/dl-alexandre/cli-tools@v0.0.1
```

## Replace version metadata

```go
import "github.com/dl-alexandre/cli-tools/version"

func Version() string { return version.String() }
```

Set values at build time:

```sh
go build -ldflags "-X github.com/dl-alexandre/cli-tools/version.Version=v1.0.0 -X github.com/dl-alexandre/cli-tools/version.BinaryName=myapp -X github.com/dl-alexandre/cli-tools/version.GitCommit=$GIT_COMMIT -X github.com/dl-alexandre/cli-tools/version.BuildTime=$BUILD_TIME"
```

## Replace update checks

```go
checker := update.New(update.Config{
	CurrentVersion: version.Version,
	BinaryName:     "myapp",
	GitHubRepo:     "owner/myapp",
	InstallCommand: "brew upgrade myapp",
})
info, err := checker.Check(false)
if err != nil {
	return err
}
return update.DisplayUpdate(info, "myapp", "table", checker.Config.InstallCommand)
```

`update.New` keeps the existing configuration shape and applies a 24-hour cache by default. Use `Check(true)` to bypass the cache.

## Add caching

```go
fileCache := cache.New(cache.DefaultDir("myapp"), 24*time.Hour)
if err := fileCache.Set("key", value, time.Hour); err != nil {
	return err
}
if value, ok := fileCache.Get("key"); ok {
	_ = value
}

memoryCache := cache.NewMemory(cache.MemoryOptions{MaxSize: 100})
```

## Add retry policy

Wrap an operation and mark only transient failures as retryable. For HTTP status responses, set `RetryAfter` using the `Retry-After` header when present.

```go
result := retry.Do(ctx, retry.Config{MaxAttempts: 3}, func(ctx context.Context) retry.Result[Response] {
	value, err := request(ctx)
	return retry.Result[Response]{Value: value, Err: err, Retry: isTransient(err)}
})
```

## Add Apple token signing

```go
signer, err := appleauth.NewFromFile(appleauth.Config{
	KeyID: "ABC123DEFG",
	TeamID: "DEF123GHIJ",
	Audience: "appstoreconnect-v1",
	TTL: 15 * time.Minute,
}, "AuthKey_ABC123DEFG.p8")
if err != nil {
	return err
}
token, err := signer.Token()
```

`Issuer` takes precedence over `TeamID` for the JWT issuer claim. The signer caches tokens and refreshes before expiration.
