package version

import (
	"strings"
	"testing"
)

func TestNormalize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"already prefixed with v", "v1.2.3", "v1.2.3"},
		{"prefixed with V", "V1.2.3", "v1.2.3"},
		{"no prefix", "1.2.3", "v1.2.3"},
		{"empty string", "", "v0.0.0"},
		{"whitespace only", "   ", "v0.0.0"},
		{"whitespace around version", "  1.2.3  ", "v1.2.3"},
		{"version with prerelease", "1.0.0-alpha", "v1.0.0-alpha"},
		{"just v", "v", "vv"},
		{"single digit", "2", "v2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Normalize(tt.input)
			if result != tt.expected {
				t.Errorf("Normalize(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		name     string
		v1       string
		v2       string
		expected int
	}{
		{"equal versions", "v1.0.0", "v1.0.0", 0},
		{"equal without prefix", "1.0.0", "1.0.0", 0},
		{"equal mixed prefix", "v1.0.0", "1.0.0", 0},
		{"major version less", "v1.0.0", "v2.0.0", -1},
		{"major version greater", "v2.0.0", "v1.0.0", 1},
		{"minor version less", "v1.1.0", "v1.2.0", -1},
		{"minor version greater", "v1.2.0", "v1.1.0", 1},
		{"patch version less", "v1.0.1", "v1.0.2", -1},
		{"patch version greater", "v1.0.2", "v1.0.1", 1},
		{"prerelease less than stable", "v1.0.0-alpha", "v1.0.0", -1},
		{"stable greater than prerelease", "v1.0.0", "v1.0.0-alpha", 1},
		{"different prereleases", "v1.0.0-alpha", "v1.0.0-beta", 0}, // Same version numbers
		{"different prereleases same version", "v1.0.0-alpha.1", "v1.0.0-alpha.2", 0},
		{"0.x versions", "v0.1.0", "v0.2.0", -1},
		{"different lengths v1 shorter", "v1.0", "v1.0.0", 0},
		{"different lengths v2 shorter", "v1.0.0", "v1.0", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Compare(tt.v1, tt.v2)
			if result != tt.expected {
				t.Errorf("Compare(%q, %q) = %d, want %d", tt.v1, tt.v2, result, tt.expected)
			}
		})
	}
}

func TestCompareOrdering(t *testing.T) {
	// Test that Compare establishes a proper ordering
	versions := []string{"v0.1.0", "v0.2.0", "v1.0.0-alpha", "v1.0.0", "v1.0.1", "v1.1.0", "v2.0.0"}

	for i := 0; i < len(versions)-1; i++ {
		for j := i + 1; j < len(versions); j++ {
			result := Compare(versions[i], versions[j])
			if result >= 0 {
				t.Errorf("Compare(%q, %q) = %d, expected < 0 (versions[%d] should be < versions[%d])",
					versions[i], versions[j], result, i, j)
			}

			// Test reverse comparison
			reverseResult := Compare(versions[j], versions[i])
			if reverseResult <= 0 {
				t.Errorf("Compare(%q, %q) = %d, expected > 0",
					versions[j], versions[i], reverseResult)
			}
		}
	}
}

func TestString(t *testing.T) {
	// Save original values
	origVersion := Version
	origBinaryName := BinaryName
	defer func() {
		Version = origVersion
		BinaryName = origBinaryName
	}()

	// Set test values
	Version = "v1.2.3"
	BinaryName = "testapp"

	result := String()
	expected := "testapp version v1.2.3"
	if result != expected {
		t.Errorf("String() = %q, want %q", result, expected)
	}
}

func TestDetailedString(t *testing.T) {
	// Save original values
	origVersion := Version
	origBinaryName := BinaryName
	origGitCommit := GitCommit
	origBuildTime := BuildTime
	defer func() {
		Version = origVersion
		BinaryName = origBinaryName
		GitCommit = origGitCommit
		BuildTime = origBuildTime
	}()

	tests := []struct {
		name         string
		version      string
		binaryName   string
		gitCommit    string
		buildTime    string
		shouldContain []string
	}{
		{
			name:         "dev build without metadata",
			version:      "dev",
			binaryName:   "myapp",
			gitCommit:    "unknown",
			buildTime:    "unknown",
			shouldContain: []string{"myapp version dev"},
		},
		{
			name:         "release build with full metadata",
			version:      "v1.2.3",
			binaryName:   "myapp",
			gitCommit:    "abc123",
			buildTime:    "2024-01-15",
			shouldContain: []string{"myapp version v1.2.3", "commit:", "abc123", "built:", "2024-01-15"},
		},
		{
			name:         "release build with only commit",
			version:      "v1.0.0",
			binaryName:   "cli",
			gitCommit:    "def456",
			buildTime:    "unknown",
			shouldContain: []string{"cli version v1.0.0", "commit:", "def456"},
		},
		{
			name:         "release build with only build time",
			version:      "v2.0.0",
			binaryName:   "tool",
			gitCommit:    "unknown",
			buildTime:    "2024-03-20",
			shouldContain: []string{"tool version v2.0.0", "built:", "2024-03-20"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Version = tt.version
			BinaryName = tt.binaryName
			GitCommit = tt.gitCommit
			BuildTime = tt.buildTime

			result := DetailedString()
			for _, expected := range tt.shouldContain {
				if !strings.Contains(result, expected) {
					t.Errorf("DetailedString() = %q, should contain %q", result, expected)
				}
			}
		})
	}
}

func TestIsDev(t *testing.T) {
	// Save original value
	origVersion := Version
	defer func() { Version = origVersion }()

	tests := []struct {
		name     string
		version  string
		expected bool
	}{
		{"dev string", "dev", true},
		{"empty string", "", true},
		{"version string", "v1.0.0", false},
		{"version without prefix", "1.0.0", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Version = tt.version
			result := IsDev()
			if result != tt.expected {
				t.Errorf("IsDev() with Version=%q = %v, want %v", tt.version, result, tt.expected)
			}
		})
	}
}

func TestFullVersion(t *testing.T) {
	// Save original value
	origVersion := Version
	defer func() { Version = origVersion }()

	tests := []struct {
		name     string
		version  string
		expected string
	}{
		{"already has prefix", "v1.2.3", "v1.2.3"},
		{"missing prefix", "1.2.3", "v1.2.3"},
		{"dev version", "dev", "vdev"},
		{"empty version", "", "v0.0.0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Version = tt.version
			result := FullVersion()
			if result != tt.expected {
				t.Errorf("FullVersion() with Version=%q = %q, want %q", tt.version, result, tt.expected)
			}
		})
	}
}

func TestGetInfo(t *testing.T) {
	// Save original values
	origVersion := Version
	origBinaryName := BinaryName
	origGitCommit := GitCommit
	origBuildTime := BuildTime
	defer func() {
		Version = origVersion
		BinaryName = origBinaryName
		GitCommit = origGitCommit
		BuildTime = origBuildTime
	}()

	// Set test values
	Version = "v2.0.0"
	BinaryName = "test-cli"
	GitCommit = "abc123def"
	BuildTime = "2024-01-20"

	info := GetInfo()

	if info.Version != "v2.0.0" {
		t.Errorf("GetInfo().Version = %q, want %q", info.Version, "v2.0.0")
	}
	if info.BinaryName != "test-cli" {
		t.Errorf("GetInfo().BinaryName = %q, want %q", info.BinaryName, "test-cli")
	}
	if info.GitCommit != "abc123def" {
		t.Errorf("GetInfo().GitCommit = %q, want %q", info.GitCommit, "abc123def")
	}
	if info.BuildTime != "2024-01-20" {
		t.Errorf("GetInfo().BuildTime = %q, want %q", info.BuildTime, "2024-01-20")
	}
}
