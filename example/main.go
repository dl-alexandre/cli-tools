// Example showing version and update helpers from cli-tools.
package main

import (
	"fmt"
	"os"

	"github.com/dl-alexandre/cli-tools/update"
	"github.com/dl-alexandre/cli-tools/version"
)

func main() {
	fmt.Println(version.DetailedString())
	if len(os.Args) < 2 || os.Args[1] != "check-update" {
		return
	}
	checker := update.New(update.Config{
		CurrentVersion: version.Version,
		BinaryName:     version.BinaryName,
		GitHubRepo:     "owner/example-cli",
		InstallCommand: "brew upgrade example-cli",
	})
	info, err := checker.Check(false)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := update.DisplayUpdate(info, version.BinaryName, "table", checker.Config.InstallCommand); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
