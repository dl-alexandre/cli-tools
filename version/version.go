// Package version provides build-time version management for Go CLI applications.
// Supports semantic versioning, git commit tracking, and build metadata.
//
// Build-time variables are typically set via ldflags:
//
//	go build -ldflags "-X github.com/dl-alexandre/cli-tools/version.Version=v1.0.0"
package version

import (
	"fmt"
	"strings"
)

// Build-time variables (set by GoReleaser or build flags via ldflags)
var (
	// Version is the current version of the CLI (e.g., "v1.2.3")
	Version = "dev"

	// BinaryName is the name of the binary (e.g., "myapp")
	BinaryName = "unknown"

	// GitCommit is the git commit hash
	GitCommit = "unknown"

	// BuildTime is the build timestamp
	BuildTime = "unknown"
)

// Info holds all version information
type Info struct {
	Version    string `json:"version"`
	BinaryName string `json:"binary_name"`
	GitCommit  string `json:"git_commit,omitempty"`
	BuildTime  string `json:"build_time,omitempty"`
}

// GetInfo returns the current version information
func GetInfo() Info {
	return Info{
		Version:    Version,
		BinaryName: BinaryName,
		GitCommit:  GitCommit,
		BuildTime:  BuildTime,
	}
}

// String returns a formatted version string
func String() string {
	return fmt.Sprintf("%s version %s", BinaryName, Version)
}

// DetailedString returns a detailed version string including commit and build time
func DetailedString() string {
	result := fmt.Sprintf("%s version %s", BinaryName, Version)

	if Version != "dev" && GitCommit != "unknown" {
		result += fmt.Sprintf("\n  commit: %s", GitCommit)
	}

	if BuildTime != "unknown" {
		result += fmt.Sprintf("\n  built:  %s", BuildTime)
	}

	return result
}

// Normalize ensures version starts with 'v' prefix
func Normalize(v string) string {
	v = strings.TrimSpace(v)
	if v == "" {
		return "v0.0.0"
	}
	if !strings.HasPrefix(v, "v") && !strings.HasPrefix(v, "V") {
		return "v" + v
	}
	return strings.ToLower(v)
}

// Compare compares two semantic versions
// Returns -1 if v1 < v2, 0 if v1 == v2, 1 if v1 > v2
func Compare(v1, v2 string) int {
	// Remove 'v' prefix
	v1 = strings.TrimPrefix(strings.ToLower(strings.TrimSpace(v1)), "v")
	v2 = strings.TrimPrefix(strings.ToLower(strings.TrimSpace(v2)), "v")

	pre1 := strings.Contains(v1, "-")
	pre2 := strings.Contains(v2, "-")
	if idx := strings.Index(v1, "-"); idx != -1 {
		v1 = v1[:idx]
	}
	if idx := strings.Index(v2, "-"); idx != -1 {
		v2 = v2[:idx]
	}

	// Split into parts
	parts1 := strings.Split(v1, ".")
	parts2 := strings.Split(v2, ".")

	// Compare each part
	maxLen := len(parts1)
	if len(parts2) > maxLen {
		maxLen = len(parts2)
	}

	for i := 0; i < maxLen; i++ {
		var num1, num2 int

		if i < len(parts1) {
			// Extract numeric part before any pre-release identifier
			part := parts1[i]
			if idx := strings.IndexAny(part, "-"); idx != -1 {
				part = part[:idx]
			}
			fmt.Sscanf(part, "%d", &num1)
		}

		if i < len(parts2) {
			part := parts2[i]
			if idx := strings.IndexAny(part, "-"); idx != -1 {
				part = part[:idx]
			}
			fmt.Sscanf(part, "%d", &num2)
		}

		if num1 < num2 {
			return -1
		}
		if num1 > num2 {
			return 1
		}
	}

	// Check pre-release status (v1.0.0-alpha < v1.0.0)
	if pre1 && !pre2 {
		return -1
	}
	if !pre1 && pre2 {
		return 1
	}

	return 0
}

// IsDev returns true if this is a development build
func IsDev() bool {
	return Version == "" || Version == "dev"
}

// FullVersion returns a normalized version string (always has v prefix)
func FullVersion() string {
	return Normalize(Version)
}
