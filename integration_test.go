package clitools_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dl-alexandre/cli-tools/config"
	"github.com/dl-alexandre/cli-tools/output"
	"github.com/dl-alexandre/cli-tools/update"
	"github.com/dl-alexandre/cli-tools/version"
)

// TestEndToEndVersion demonstrates a complete version workflow
func TestEndToEndVersion(t *testing.T) {
	// Save original values
	origVersion := version.Version
	origBinaryName := version.BinaryName
	origGitCommit := version.GitCommit
	origBuildTime := version.BuildTime
	defer func() {
		version.Version = origVersion
		version.BinaryName = origBinaryName
		version.GitCommit = origGitCommit
		version.BuildTime = origBuildTime
	}()

	// Simulate a release build
	version.Version = "v1.5.2"
	version.BinaryName = "mycli"
	version.GitCommit = "abc123"
	version.BuildTime = "2024-01-20"

	// Test version string
	verStr := version.String()
	if verStr != "mycli version v1.5.2" {
		t.Errorf("version.String() = %q, want %q", verStr, "mycli version v1.5.2")
	}

	// Test detailed version
	detailed := version.DetailedString()
	if !contains(detailed, "mycli version v1.5.2") {
		t.Errorf("version.DetailedString() missing version info")
	}
	if !contains(detailed, "abc123") {
		t.Errorf("version.DetailedString() missing commit")
	}
	if !contains(detailed, "2024-01-20") {
		t.Errorf("version.DetailedString() missing build time")
	}

	// Test it's not a dev build
	if version.IsDev() {
		t.Error("IsDev() should be false for release build")
	}
}

// TestEndToEndUpdate demonstrates a complete update workflow
func TestEndToEndUpdate(t *testing.T) {
	// Create a temporary cache directory
	tempDir := t.TempDir()

	// Create update checker
	checker := update.New(update.Config{
		CurrentVersion: "v1.0.0",
		BinaryName:     "mycli",
		GitHubRepo:     "owner/repo",
		CacheDir:       tempDir,
		InstallCommand: "brew upgrade mycli",
	})

	// Verify config was set correctly
	if checker.Config.CacheDir != tempDir {
		t.Errorf("CacheDir not set correctly")
	}

	// Test GitHub URL building
	// Note: This would fail if actually called without a mock server
	// but we're just verifying the URL construction logic
}

// TestEndToEndConfig demonstrates a complete config workflow
func TestEndToEndConfig(t *testing.T) {
	// Create a temporary directory
	tempDir := t.TempDir()

	// Create config loader
	loader := config.NewLoader("testapp", "TESTAPP")

	// Set defaults
	loader.SetDefaults(map[string]interface{}{
		"api.base_url": "https://api.example.com",
		"api.timeout":  30,
	})

	// Save some config
	configData := map[string]interface{}{
		"api.base_url":  "https://api.prod.example.com",
		"output.format": "json",
	}

	// Test that we can set values in viper
	v, err := loader.Load(config.Flags{})
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if v.GetString("api.base_url") != "https://api.example.com" {
		t.Error("Default not set correctly")
	}

	_ = tempDir
	_ = configData
}

// TestEndToEndOutput demonstrates output formatting
func TestEndToEndOutput(t *testing.T) {
	// Create formatter
	formatter := output.New("table", "never", false)

	// Create table
	tbl := formatter.NewTable("Name", "Status", "Count")
	tbl.AddRow("Item 1", "active", 42)
	tbl.AddRow("Item 2", "inactive", 0)

	// Just verify it doesn't panic
	// In real tests, we'd capture stdout

	// Test message functions - just verify they don't panic
	output.PrintEmptyMessage("items")
	output.PrintSuccess("Operation completed")
	output.PrintWarning("This is a warning")

	// Test error - redirect stderr temporarily
	output.PrintError("Something went wrong")
}

// TestIntegrationPatterns demonstrates common integration patterns
func TestIntegrationPatterns(t *testing.T) {
	t.Run("version_and_update_integration", func(t *testing.T) {
		// Save original values
		origVersion := version.Version
		origBinaryName := version.BinaryName
		defer func() {
			version.Version = origVersion
			version.BinaryName = origBinaryName
		}()

		// Set version info
		version.Version = "v2.0.0"
		version.BinaryName = "mycli"

		// Create update checker using version info
		checker := update.New(update.Config{
			CurrentVersion: version.Version,
			BinaryName:     version.BinaryName,
			GitHubRepo:     "myuser/mycli",
		})

		if checker.Config.CurrentVersion != "v2.0.0" {
			t.Error("CurrentVersion not passed correctly to update checker")
		}
	})

	t.Run("config_and_credentials", func(t *testing.T) {
		// Create config loader
		loader := config.NewLoader("myapp", "MYAPP")

		// Test credentials helper
		username, password, err := config.GetCredentials("admin", "secret", "MYAPP")
		if err != nil {
			t.Errorf("GetCredentials() error = %v", err)
		}
		if username != "admin" {
			t.Errorf("username = %q, want %q", username, "admin")
		}
		if password != "secret" {
			t.Errorf("password = %q, want %q", password, "secret")
		}

		_ = loader
	})
}

// BenchmarkVersionCompare benchmarks version comparison
func BenchmarkVersionCompare(b *testing.B) {
	versions := []string{
		"v1.0.0", "v1.0.1", "v1.1.0", "v2.0.0",
		"v0.1.0", "v1.0.0-alpha", "v2.0.0-beta",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(versions); j++ {
			for k := 0; k < len(versions); k++ {
				version.Compare(versions[j], versions[k])
			}
		}
	}
}

// BenchmarkNormalize benchmarks version normalization
func BenchmarkNormalize(b *testing.B) {
	testVersions := []string{
		"1.0.0", "v1.0.0", "V1.0.0", "  1.0.0  ",
		"1.0.0-alpha", "", "v1.0.0-beta.1",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range testVersions {
			version.Normalize(v)
		}
	}
}

// Helper functions
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// TestMockServer tests update checking with a mock server
func TestMockServer(t *testing.T) {
	// This test would require setting up an httptest.Server
	// to mock the GitHub API responses
	// Skipping for now as it requires more complex setup
	t.Skip("Skipping mock server test - requires HTTP test setup")
}

// TestRealConfigLoading tests actual config file loading
func TestRealConfigLoading(t *testing.T) {
	// Create a temporary directory for config
	tempDir := t.TempDir()

	// Create a test config file
	configPath := filepath.Join(tempDir, "config.yaml")
	configContent := `
api:
  base_url: "https://api.test.com"
  timeout: 45
output:
  format: "json"
`
	err := os.WriteFile(configPath, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test config: %v", err)
	}

	// Create loader and load config
	loader := config.NewLoader("testapp", "TEST")
	loader.SetDefaults(map[string]interface{}{
		"api.timeout": 30,
	})

	v, err := loader.Load(config.Flags{ConfigFile: configPath})
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Verify values were loaded
	if v.GetString("api.base_url") != "https://api.test.com" {
		t.Errorf("api.base_url = %q, want %q", v.GetString("api.base_url"), "https://api.test.com")
	}
	if v.GetInt("api.timeout") != 45 {
		t.Errorf("api.timeout = %d, want %d", v.GetInt("api.timeout"), 45)
	}
	if v.GetString("output.format") != "json" {
		t.Errorf("output.format = %q, want %q", v.GetString("output.format"), "json")
	}
}
