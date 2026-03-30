# cli-tools

[![CI](https://github.com/dl-alexandre/cli-tools/actions/workflows/ci.yml/badge.svg)](https://github.com/dl-alexandre/cli-tools/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dl-alexandre/cli-tools)](https://goreportcard.com/report/github.com/dl-alexandre/cli-tools)
[![GoDoc](https://godoc.org/github.com/dl-alexandre/cli-tools?status.svg)](https://godoc.org/github.com/dl-alexandre/cli-tools)

A Go library providing common functionality for CLI applications. Reduce boilerplate code and ensure consistent behavior across your CLI tools.

## Features

- **Version Management** - Semantic versioning, build-time variables, version comparison
- **Update Checking** - Automatic GitHub release checking with caching
- **Output Formatting** - Table and JSON output, color detection, message helpers
- **Configuration** - Viper-based config with environment variable support
- **CLI Framework Helpers** - Extensions for the Kong CLI framework

## Installation

```bash
go get github.com/dl-alexandre/cli-tools
```

## Quick Start

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

Build with version info:
```bash
go build -ldflags "-X github.com/dl-alexandre/cli-tools/version.Version=v1.0.0"
```

## Modules

### `version`
Build-time version management including:
- Semantic version handling
- Version comparison
- Development build detection
- Formatted version string output

### `update`
GitHub release update checking:
- Automatic update detection
- 24-hour caching
- CI environment detection (skips in CI)
- Standardized update notification format

### `output`
Output formatting utilities:
- Table formatting with `rodaine/table`
- JSON output formatting
- Color/auto-detection with `go-isatty`
- Common message helpers (success, error, warning)
- Banner printing

### `config`
Configuration management:
- Viper-based configuration
- Standard file locations (`~/.config/{binary}/`)
- Environment variable binding
- Config precedence: flags > env > file > defaults
- Credential helper for username/password patterns

### `kongx`
Kong CLI framework extensions:
- Common flag definitions
- Standard Run method patterns
- Help formatting

## Usage

### In Your CLI

```go
// In your go.mod
require github.com/dl-alexandre/cli-tools v1.0.0

// In your main.go or version.go
import "github.com/dl-alexandre/cli-tools/version"

// Set build-time variables via ldflags:
// go build -ldflags "-X github.com/dl-alexandre/cli-tools/version.Version=v1.0.0 -X github.com/dl-alexandre/cli-tools/version.BinaryName=myapp"

// Use in your CLI:
fmt.Println(version.String())
```

### Example: Update Checking

```go
import (
    "github.com/dl-alexandre/cli-tools/update"
    "github.com/dl-alexandre/cli-tools/version"
)

// Create checker
checker := update.New(update.Config{
    CurrentVersion: version.Version,
    BinaryName:     "myapp",
    GitHubRepo:     "myuser/myapp",  // or GitHubOwner: "myuser", GitHubRepoName: "myapp"
    InstallCommand: "brew upgrade myapp",  // optional
})

// Check for updates (uses cache)
info, err := checker.Check(false)
if err != nil {
    return err
}

// Display result
update.DisplayUpdate(info, "myapp", "table", "brew upgrade myapp")
```

### Example: Configuration

```go
import "github.com/dl-alexandre/cli-tools/config"

// Create loader
loader := config.NewLoader("myapp", "MYAPP")

// Set defaults
loader.SetDefaults(map[string]interface{}{
    "api.base_url": "https://api.example.com",
    "api.timeout": 30,
})

// Load config
v, err := loader.Load(config.Flags{
    ConfigFile: cfg.ConfigFile,
})

// Get credentials (common pattern)
username, password, err := config.GetCredentials(
    flags.Username, flags.Password, "MYAPP",
)
```

## Development

### Local Development

Use a `replace` directive in your CLI's `go.mod` for local development:

```go
// In your CLI's go.mod
replace github.com/dl-alexandre/cli-tools => ../cli-tools
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with race detection
go test -race ./...

# Run benchmarks
go test -bench=. ./...

# Run specific package tests
go test ./version
go test ./update
go test ./output
go test ./config
```

### Test Coverage

The library includes comprehensive test coverage:
- **Unit tests** for each module
- **Integration tests** demonstrating end-to-end workflows
- **Benchmarks** for performance-critical operations

### Writing Tests

When contributing, ensure your tests:
1. Cover both success and error paths
2. Use table-driven tests for multiple test cases
3. Clean up resources (use `t.TempDir()` for temp files)
4. Include benchmarks for performance-sensitive code

## Versioning

This library follows [Semantic Versioning](https://semver.org/):
- **v0.x.x** - Initial development, API may change
- **v1.x.x** - Stable API with backward compatibility guarantees

Current version: **v0.0.0** - Initial release for testing with dl-alexandre CLI tools

## Releasing

To create a new release:

```bash
# Tag a new version
git tag v0.1.0
git push origin v0.1.0
```

The [release workflow](.github/workflows/release.yml) will automatically create a GitHub release.

## License

MIT License - See [LICENSE](LICENSE) file for details

## Contributing

Contributions should:
1. Follow existing patterns and conventions
2. Be useful across multiple CLI tools
3. Not introduce breaking changes without versioning
4. Include tests for new functionality
