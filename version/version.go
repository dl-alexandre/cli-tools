// Package version provides build-time version management for Go CLI applications.
// Supports semantic versioning, git commit tracking, and build metadata.
//
// Build-time variables are typically set via ldflags:
//
//	go build -ldflags "-X github.com/dl-alexandre/cli-tools/version.Version=v1.0.0"
package version

import (
	"fmt"
	"strconv"
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
	major1, minor1, patch1, pre1 := parseVersion(v1)
	major2, minor2, patch2, pre2 := parseVersion(v2)
	for _, pair := range [][2]uint64{{major1, major2}, {minor1, minor2}, {patch1, patch2}} {
		if pair[0] < pair[1] {
			return -1
		}
		if pair[0] > pair[1] {
			return 1
		}
	}
	return comparePrerelease(pre1, pre2)
}

func parseVersion(value string) (major, minor, patch uint64, prerelease string) {
	value = strings.TrimPrefix(strings.ToLower(strings.TrimSpace(value)), "v")
	if before, _, ok := strings.Cut(value, "+"); ok {
		value = before // build metadata does not affect precedence
	}
	if before, after, ok := strings.Cut(value, "-"); ok {
		value, prerelease = before, after
	}
	parts := strings.Split(value, ".")
	numbers := []*uint64{&major, &minor, &patch}
	for i := 0; i < min(len(parts), len(numbers)); i++ {
		*numbers[i], _ = strconv.ParseUint(parts[i], 10, 64)
	}
	return major, minor, patch, prerelease
}

func comparePrerelease(a, b string) int {
	if a == b {
		return 0
	}
	if a == "" {
		return 1
	}
	if b == "" {
		return -1
	}
	left, right := strings.Split(a, "."), strings.Split(b, ".")
	for i := 0; i < min(len(left), len(right)); i++ {
		if left[i] == right[i] {
			continue
		}
		leftNumber, leftErr := strconv.ParseUint(left[i], 10, 64)
		rightNumber, rightErr := strconv.ParseUint(right[i], 10, 64)
		switch {
		case leftErr == nil && rightErr == nil:
			if leftNumber < rightNumber {
				return -1
			}
			return 1
		case leftErr == nil:
			return -1
		case rightErr == nil:
			return 1
		case left[i] < right[i]:
			return -1
		default:
			return 1
		}
	}
	if len(left) < len(right) {
		return -1
	}
	return 1
}

// IsDev returns true if this is a development build
func IsDev() bool {
	return Version == "" || Version == "dev"
}

// FullVersion returns a normalized version string (always has v prefix)
func FullVersion() string {
	return Normalize(Version)
}
