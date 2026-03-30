# Migration Guide

This guide shows how to migrate an existing CLI to use the `cli-tools` shared library.

## Step 1: Update go.mod

Add the dependency:

```bash
go get github.com/dl-alexandre/cli-tools
```

Or manually add to `go.mod`:

```go
require github.com/dl-alexandre/cli-tools v1.0.0
```

## Step 2: Replace version.go

**Before (internal/cli/version.go):**

```go
package cli

var (
    Version    = "dev"
    BinaryName = "unifi"
    GitHubRepo = "Local-UniFi-CLI"
    GitCommit  = "unknown"
    BuildTime  = "unknown"
)
```

**After:**

```go
package cli

import "github.com/dl-alexandre/cli-tools/version"

func init() {
    // Set your CLI-specific values
    version.BinaryName = "unifi"
    version.GitHubRepo = "Local-UniFi-CLI"
    // Version, GitCommit, BuildTime are set at build time via ldflags
}

// Re-export for backward compatibility
var (
    Version   = version.Version
    GitCommit = version.GitCommit
    BuildTime = version.BuildTime
)
```

**Update your build flags:**

Change from:
```bash
go build -ldflags "-X main.Version=1.0.0"
```

To:
```bash
go build -ldflags "-X github.com/dl-alexandre/cli-tools/version.Version=1.0.0"
```

## Step 3: Replace update.go

**Before (internal/cli/update.go):**

~300 lines of duplicated code for GitHub API calls, caching, version comparison...

**After:**

```go
package cli

import (
    "github.com/dl-alexandre/cli-tools/update"
    "github.com/dl-alexandre/cli-tools/version"
)

type UpdateCheckCmd struct {
    Force  bool   `help:"Force check, bypassing cache" flag:"force"`
    Format string `help:"Output format" enum:"table,json" default:"table"`
}

func (c *UpdateCheckCmd) Run(globals *Globals) error {
    checker := update.New(update.Config{
        CurrentVersion: version.Version,
        BinaryName:     version.BinaryName,
        GitHubRepo:     "myuser/myapp",  // full repo path
        InstallCommand: "brew upgrade myapp",  // optional
    })

    info, err := checker.Check(c.Force)
    if err != nil {
        return err
    }

    return update.DisplayUpdate(info, version.BinaryName, c.Format, "brew upgrade myapp")
}

// Auto-update check at startup:
func AutoUpdateCheck() {
    checker := update.New(update.Config{
        CurrentVersion: version.Version,
        BinaryName:     version.BinaryName,
        GitHubRepo:     "myuser/myapp",
        InstallCommand: "brew upgrade myapp",
    })
    checker.AutoCheck()
}
```

**Lines of code: 300 → 30 (90% reduction)**

## Step 4: Replace output/formatter.go

**Before (internal/pkg/output/formatter.go):**

~477 lines with table formatting, color detection, JSON output...

**After:**

```go
package output

import (
    "github.com/dl-alexandre/cli-tools/output"
)

// Use the shared formatter
type Formatter = output.Formatter

// Custom data types for your CLI
type SiteData struct {
    ID          string
    Name        string
    Description string
    Devices     int
    Clients     int
}

func PrintSitesTable(sites []SiteData, formatter *output.Formatter) {
    if len(sites) == 0 {
        output.PrintEmptyMessage("sites")
        return
    }

    tbl := formatter.NewTable("ID", "Name", "Description", "Devices", "Clients")
    for _, site := range sites {
        desc := site.Description
        if desc == "" {
            desc = "-"
        }
        tbl.AddRow(site.ID, site.Name, desc, site.Devices, site.Clients)
    }
    tbl.Print()
}
```

## Step 5: Replace config/config.go

**Before (internal/pkg/config/config.go):**

~220 lines of viper setup, env binding, config file handling...

**After:**

```go
package config

import (
    "github.com/dl-alexandre/cli-tools/config"
)

type Config struct {
    API    APIConfig    `mapstructure:"api"`
    Auth   AuthConfig   `mapstructure:"auth"`
    Output OutputConfig `mapstructure:"output"`
}

type APIConfig struct {
    BaseURL string `mapstructure:"base_url"`
    Timeout int    `mapstructure:"timeout"`
}

type AuthConfig struct {
    Username string `mapstructure:"username"`
}

type OutputConfig struct {
    Format    string `mapstructure:"format"`
    Color     string `mapstructure:"color"`
    NoHeaders bool   `mapstructure:"no_headers"`
}

func Load(flags GlobalFlags) (*Config, error) {
    loader := config.NewLoader("unifi", "UNIFI")
    
    loader.SetDefaults(map[string]interface{}{
        "api.base_url": "https://unifi.local",
        "api.timeout": 30,
        "output.format": "table",
        "output.color": "auto",
    })

    v, err := loader.Load(config.Flags{
        ConfigFile: flags.ConfigFile,
    })
    if err != nil {
        return nil, err
    }

    var cfg Config
    if err := v.Unmarshal(&cfg); err != nil {
        return nil, err
    }

    // Apply CLI flags (override config)
    if flags.BaseURL != "" {
        cfg.API.BaseURL = flags.BaseURL
    }
    if flags.Timeout > 0 {
        cfg.API.Timeout = flags.Timeout
    }

    return &cfg, nil
}

// Use shared credential helper
func GetCredentials(flagsUsername, flagsPassword string) (string, string, error) {
    return config.GetCredentials(flagsUsername, flagsPassword, "UNIFI")
}
```

## Step 6: Update main.go

**Add auto-update check at startup:**

```go
func main() {
    // ... existing setup ...

    // Auto-check for updates (non-blocking)
    checker := update.New(update.Config{
        CurrentVersion: version.Version,
        BinaryName:     version.BinaryName,
        GitHubRepo:     version.GitHubRepo,
    })
    checker.AutoCheck()

    // Run command
    err = ctx.Run(&cli.Globals)
    // ...
}
```

## Step 7: Update .goreleaser.yml

**Update ldflags to point to cli-tools:**

```yaml
builds:
  - id: mycli
    binary: mycli
    ldflags:
      - -s -w
      - -X github.com/dl-alexandre/cli-tools/version.Version={{.Version}}
      - -X github.com/dl-alexandre/cli-tools/version.GitCommit={{.Commit}}
      - -X github.com/dl-alexandre/cli-tools/version.BuildTime={{.Date}}
      - -X github.com/dl-alexandre/cli-tools/version.BinaryName=mycli
      - -X github.com/dl-alexandre/cli-tools/version.GitHubRepo=mycli-repo
```

## Benefits Summary

| Metric | Before | After | Improvement |
|--------|--------|-------|-------------|
| version.go | 19 lines | 12 lines | 37% reduction |
| update.go | ~300 lines | ~30 lines | 90% reduction |
| output/formatter.go | ~477 lines | ~50 lines | 89% reduction |
| config/config.go | ~220 lines | ~70 lines | 68% reduction |
| **Total** | **~1,016 lines** | **~162 lines** | **84% reduction** |

## Testing After Migration

1. Build the CLI:
   ```bash
   go build -ldflags "-X github.com/dl-alexandre/cli-tools/version.Version=v1.0.0 -X github.com/dl-alexandre/cli-tools/version.BinaryName=mycli"
   ```

2. Test version command:
   ```bash
   ./mycli version
   ```

3. Test update check:
   ```bash
   ./mycli check-update --force
   ```

4. Test table output:
   ```bash
   ./mycli sites list
   ./mycli sites list --format json
   ```

5. Test config loading:
   ```bash
   ./mycli sites list --config /path/to/config.yaml
   ```

## Troubleshooting

### Issue: Version shows "dev" after build
**Solution:** Ensure ldflags point to `github.com/dl-alexandre/cli-tools/version.Version`, not the old local path.

### Issue: Import cycle errors
**Solution:** Make sure you're not importing `cli-tools` from your main package back into itself. Use type aliases where needed.

### Issue: Build fails with "undefined: version.BinaryName"
**Solution:** The `BinaryName` and `GitHubRepo` must be set in your CLI code before the shared library can use them. See Step 2.
