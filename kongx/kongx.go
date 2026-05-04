// Package kongx provides Kong CLI framework extensions.
// Common patterns and helpers for CLIs built with the Kong framework.
//
// Note: This package is specifically for Kong-based CLIs.
// For Cobra-based CLIs, use similar patterns in your own code.
package kongx

import (
	"fmt"

	"github.com/alecthomas/kong"
)

// Context holds common context for Kong-based CLIs
type Context struct {
	// Config is the loaded configuration (type depends on CLI)
	Config interface{}

	// Debug enables debug output
	Debug bool
}

// CommonFlags defines standard flags for Kong-based CLIs
type CommonFlags struct {
	ConfigFile string `help:"Config file path" short:"c" env:"CONFIG_FILE"`
	Format     string `help:"Output format" enum:"table,json" default:"table"`
	Color      string `help:"Color output" enum:"auto,always,never" default:"auto"`
	NoHeaders  bool   `help:"Disable table headers"`
	Debug      bool   `help:"Enable debug output"`
	Version    kong.VersionFlag `help:"Show version" short:"v"`
}

// CommonUpdateFlags defines standard update-related flags
type CommonUpdateFlags struct {
	Force  bool   `help:"Force check, bypassing cache" flag:"force"`
	Format string `help:"Output format" enum:"table,json" default:"table"`
}

// RunFunc is the standard Run method signature for Kong commands
type RunFunc func(ctx *Context) error

// ErrorHandler provides consistent error formatting
func ErrorHandler(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("✗ %w", err)
}

// HelpFormatter returns a custom help formatter
func HelpFormatter(appName, description string) string {
	return fmt.Sprintf(`%s - %s

Usage:
  %s <command> [flags]

Flags:
{{flags .Flags}}

Commands:
{{commands .Commands}}
`, appName, description, appName)
}

// Parse parses command line arguments using Kong with standard configuration
// app: the application struct with commands and flags
// appName: binary name (e.g., "unifi", "cimis")
// description: short description of the CLI
// version: version string for --version flag
func Parse(app interface{}, appName, description, version string, args []string) (*kong.Context, error) {
	parser, err := kong.New(
		app,
		kong.Name(appName),
		kong.Description(description),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Tree:    true,
		}),
		kong.Vars{
			"version": version,
		},
	)
	if err != nil {
		return nil, err
	}

	ctx, err := parser.Parse(args)
	return ctx, err
}
