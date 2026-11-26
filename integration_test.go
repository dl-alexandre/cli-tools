package clitools_test

import (
	"testing"

	"github.com/dl-alexandre/cli-tools/update"
	"github.com/dl-alexandre/cli-tools/version"
)

func TestVersionAndUpdateConfiguration(t *testing.T) {
	originalVersion, originalBinaryName := version.Version, version.BinaryName
	version.Version, version.BinaryName = "v2.0.0", "mycli"
	t.Cleanup(func() {
		version.Version, version.BinaryName = originalVersion, originalBinaryName
	})

	if got, want := version.String(), "mycli version v2.0.0"; got != want {
		t.Fatalf("version.String() = %q, want %q", got, want)
	}
	checker := update.New(update.Config{
		CurrentVersion: version.Version,
		BinaryName:     version.BinaryName,
		GitHubRepo:     "owner/mycli",
	})
	if checker.Config.CurrentVersion != version.Version {
		t.Fatalf("update checker version = %q, want %q", checker.Config.CurrentVersion, version.Version)
	}
}
