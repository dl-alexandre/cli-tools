// Example showing how a CLI would use cli-tools
// This demonstrates the before/after migration pattern

package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/dl-alexandre/cli-tools/config"
	"github.com/dl-alexandre/cli-tools/output"
	"github.com/dl-alexandre/cli-tools/update"
	"github.com/dl-alexandre/cli-tools/version"
)

// CLI represents the command-line interface
type CLI struct {
	Globals

	Sites    SitesCmd    `cmd:"" help:"Manage sites"`
	Devices  DevicesCmd  `cmd:"" help:"Manage devices"`
	Check    CheckCmd    `cmd:"check-update" help:"Check for updates"`
	Version  VersionCmd  `cmd:"version" help:"Show version information"`
}

// Globals holds global configuration
type Globals struct {
	ConfigFile string `help:"Config file path" short:"c"`
	Format     string `help:"Output format" enum:"table,json" default:"table"`
	Color      string `help:"Color output" enum:"auto,always,never" default:"auto"`
	NoHeaders  bool   `help:"Disable table headers"`
	Debug      bool   `help:"Enable debug output"`
	Version    kong.VersionFlag `help:"Show version" short:"v"`
}

// SitesCmd manages UniFi sites
type SitesCmd struct {
	List ListSitesCmd `cmd:"" help:"List all sites"`
}

type ListSitesCmd struct{}

func (c *ListSitesCmd) Run(globals *Globals) error {
	// Use shared output formatter
	formatter := output.New(globals.Format, globals.Color, globals.NoHeaders)

	// Example data
	sites := []map[string]interface{}{
		{"id": "default", "name": "Default", "devices": 5, "clients": 12},
		{"id": "home", "name": "Home", "devices": 3, "clients": 8},
	}

	if globals.Format == "json" {
		return formatter.PrintJSON(sites)
	}

	// Build table using shared table builder
	tbl := formatter.NewTable("ID", "Name", "Devices", "Clients")
	for _, site := range sites {
		tbl.AddRow(site["id"], site["name"], site["devices"], site["clients"])
	}
	tbl.Print()

	return nil
}

// DevicesCmd manages UniFi devices
type DevicesCmd struct {
	List ListDevicesCmd `cmd:"" help:"List all devices"`
}

type ListDevicesCmd struct{}

func (c *ListDevicesCmd) Run(globals *Globals) error {
	formatter := output.New(globals.Format, globals.Color, globals.NoHeaders)

	devices := []map[string]interface{}{
		{"mac": "aa:bb:cc:dd:ee:ff", "name": "AP-Living-Room", "model": "U6-Pro", "status": "online"},
	}

	if globals.Format == "json" {
		return formatter.PrintJSON(devices)
	}

	tbl := formatter.NewTable("MAC", "Name", "Model", "Status")
	for _, dev := range devices {
		tbl.AddRow(dev["mac"], dev["name"], dev["model"], dev["status"])
	}
	tbl.Print()

	return nil
}

// CheckCmd checks for updates
type CheckCmd struct {
	Force  bool   `help:"Force check, bypassing cache" flag:"force"`
	Format string `help:"Output format" enum:"table,json" default:"table"`
}

func (c *CheckCmd) Run(globals *Globals) error {
	// Use shared update checker
	checker := update.New(update.Config{
		CurrentVersion: version.Version,
		BinaryName:     version.BinaryName,
		GitHubRepo:     version.GitHubRepo,
	})

	info, err := checker.Check(c.Force)
	if err != nil {
		return err
	}

	return update.DisplayUpdate(info, version.BinaryName, c.Format)
}

// VersionCmd shows version information
type VersionCmd struct {
	CheckLatest bool `help:"Check for latest version"`
}

func (c *VersionCmd) Run(globals *Globals) error {
	// Use shared version formatting
	fmt.Println(version.DetailedString())

	if c.CheckLatest {
		fmt.Println("\nChecking for updates...")
		checker := update.New(update.Config{
			CurrentVersion: version.Version,
			BinaryName:     version.BinaryName,
			GitHubRepo:     "myuser/example-cli",
		})

		info, err := checker.Check(false)
		if err != nil {
			return err
		}

		if info.UpdateAvailable {
			fmt.Printf("Update available: %s\n", info.LatestVersion)
		} else {
			fmt.Println("You are running the latest version.")
		}
	}

	return nil
}

func main() {
	var cli CLI

	parser := kong.Must(&cli,
		kong.Name(version.BinaryName),
		kong.Description("Example CLI using cli-tools shared library"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			"version": version.Version,
		},
	)

	ctx, err := parser.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "✗ %v\n", err)
		os.Exit(1)
	}

	// Auto-check for updates in background (non-blocking)
	checker := update.New(update.Config{
		CurrentVersion: version.Version,
		BinaryName:     version.BinaryName,
		GitHubRepo:     "myuser/example-cli",
	})
	checker.AutoCheck()

	// Run the command
	err = ctx.Run(&cli.Globals)
	if err != nil {
		fmt.Fprintf(os.Stderr, "✗ %v\n", err)
		os.Exit(1)
	}
}
