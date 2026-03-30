// Package update provides GitHub release update checking for Go CLI applications.
// Supports checking GitHub releases, caching results, and displaying update notifications.
//
// Basic usage:
//   checker := update.New(update.Config{
//       CurrentVersion: "v1.0.0",
//       GitHubRepo:     "owner/repo",
//   })
//   info, err := checker.Check(false)
package update

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	cacheKey = "update_check"
	cacheTTL = 24 * time.Hour // Check once per day
)

// GitHubRelease represents the release information from GitHub API
type GitHubRelease struct {
	TagName     string    `json:"tag_name"`
	Name        string    `json:"name"`
	PublishedAt time.Time `json:"published_at"`
	HTMLURL     string    `json:"html_url"`
	Body        string    `json:"body"`
	Prerelease  bool      `json:"prerelease"`
	Draft       bool      `json:"draft"`
}

// Info holds the update check result
type Info struct {
	CurrentVersion  string `json:"current_version"`
	LatestVersion   string `json:"latest_version"`
	UpdateAvailable bool   `json:"update_available"`
	ReleaseURL      string `json:"release_url"`
	PublishedAt     string `json:"published_at"`
	IsPrerelease    bool   `json:"is_prerelease"`
}

// Config holds configuration for update checking
type Config struct {
	// CurrentVersion is the version of the running binary
	CurrentVersion string

	// BinaryName is the name of the binary (for User-Agent and messages)
	BinaryName string

	// GitHubRepo is the GitHub repository in "owner/repo" format (e.g., "myuser/myapp")
	// Alternatively, can use GitHubOwner + GitHubRepoName separately
	GitHubRepo string

	// GitHubOwner is the GitHub username/organization (e.g., "myuser")
	GitHubOwner string

	// GitHubRepoName is the repository name (e.g., "myapp")
	GitHubRepoName string

	// CacheDir is the directory to store cached update info (default: ~/.cache/{binary})
	CacheDir string

	// CacheTTL is how long to cache update checks (default: 24 hours)
	CacheTTL time.Duration

	// InstallCommand is shown to users when an update is available
	// (e.g., "brew upgrade myapp" or "go install github.com/user/repo@latest")
	InstallCommand string
}

// Checker handles update checking
type Checker struct {
	config Config
}

// New creates a new update checker with the given configuration
func New(config Config) *Checker {
	if config.CacheTTL == 0 {
		config.CacheTTL = cacheTTL
	}
	if config.CacheDir == "" && config.BinaryName != "" {
		config.CacheDir = defaultCacheDir(config.BinaryName)
	}
	return &Checker{config: config}
}

// Check checks for available updates
func (c *Checker) Check(force bool) (*Info, error) {
	currentVersion := c.config.CurrentVersion
	if currentVersion == "" || currentVersion == "dev" {
		currentVersion = "v0.0.0"
	}

	// Try to get from cache first
	if !force && c.config.CacheDir != "" {
		if cached, ok := c.getCached(); ok {
			return cached, nil
		}
	}

	// Fetch latest release from GitHub
	info, err := c.fetchLatest(currentVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to check for updates: %w", err)
	}

	// Cache the result
	if c.config.CacheDir != "" {
		c.cacheResult(info)
	}

	return info, nil
}

// fetchLatest queries GitHub API for the latest release
func (c *Checker) fetchLatest(currentVersion string) (*Info, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	apiURL := c.buildGitHubAPIURL()

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	// GitHub API requires a User-Agent header
	req.Header.Set("User-Agent", fmt.Sprintf("%s/%s", c.config.BinaryName, currentVersion))
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitHub API returned %d: %s", resp.StatusCode, string(body))
	}

	var release GitHubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}

	// Normalize version strings
	latestVersion := normalizeVersion(release.TagName)
	currentNormalized := normalizeVersion(currentVersion)

	// Compare versions
	updateAvailable := compareVersions(currentNormalized, latestVersion) < 0

	return &Info{
		CurrentVersion:  currentVersion,
		LatestVersion:   latestVersion,
		UpdateAvailable: updateAvailable,
		ReleaseURL:      release.HTMLURL,
		PublishedAt:     release.PublishedAt.Format("2006-01-02"),
		IsPrerelease:    release.Prerelease,
	}, nil
}

// buildGitHubAPIURL constructs the GitHub API URL
func (c *Checker) buildGitHubAPIURL() string {
	owner := c.config.GitHubOwner
	repo := c.config.GitHubRepoName

	// If GitHubRepo is provided in "owner/repo" format, parse it
	if c.config.GitHubRepo != "" {
		parts := strings.Split(c.config.GitHubRepo, "/")
		if len(parts) == 2 {
			owner = parts[0]
			repo = parts[1]
		}
	}

	// Fall back to binary name if repo not specified
	if repo == "" {
		repo = c.config.BinaryName
	}

	return fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)
}

// AutoCheck performs a background update check (for use at startup)
// It returns immediately and doesn't block
func (c *Checker) AutoCheck() {
	// Skip in CI environments
	if isCIEnvironment() {
		return
	}

	// Check if we've already checked recently
	if _, ok := c.getCached(); ok {
		return
	}

	// Perform check in background
	go func() {
		currentVersion := c.config.CurrentVersion
		if currentVersion == "" || currentVersion == "dev" {
			currentVersion = "v0.0.0"
		}

		info, err := c.fetchLatest(currentVersion)
		if err != nil {
			return // Silently fail on auto-check
		}

		// Cache the result
		c.cacheResult(info)

		// Only print if update is available
		if info.UpdateAvailable {
			fmt.Println()
			fmt.Printf("📦 A new version is available: %s (current: %s)\n", info.LatestVersion, info.CurrentVersion)
			if c.config.InstallCommand != "" {
				fmt.Printf("   Run '%s check-update' for details or upgrade with: %s\n", c.config.BinaryName, c.config.InstallCommand)
			} else {
				fmt.Printf("   Run '%s check-update' for details\n", c.config.BinaryName)
			}
			fmt.Println()
		}
	}()
}

// DisplayUpdate shows update information in the standard dl-alexandre format
func DisplayUpdate(info *Info, binaryName string, format string) error {
	switch format {
	case "json":
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(info)
	default:
		if info.UpdateAvailable {
			// Print notification banner
			fmt.Println()
			fmt.Println("╔══════════════════════════════════════════════════════════════╗")
			fmt.Println("║                    UPDATE AVAILABLE                            ║")
			fmt.Println("╚══════════════════════════════════════════════════════════════╝")
			fmt.Println()
			fmt.Printf("Current version: %s\n", info.CurrentVersion)
			fmt.Printf("Latest version:  %s\n", info.LatestVersion)
			fmt.Printf("Published:       %s\n", info.PublishedAt)
			fmt.Println()
			fmt.Println("Install the latest version:")
			fmt.Printf("  brew upgrade %s\n", binaryName)
			fmt.Println()
			fmt.Printf("Or download from: %s\n", info.ReleaseURL)
			fmt.Println()

			if info.IsPrerelease {
				fmt.Println("⚠️  This is a pre-release version.")
				fmt.Println()
			}
		} else {
			fmt.Printf("✓ You're running the latest version (%s)\n", info.CurrentVersion)
		}
	}
	return nil
}

// cacheEntry represents a cached update check
type cacheEntry struct {
	Data      *Info     `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}

// cacheResult saves update info to cache
func (c *Checker) cacheResult(info *Info) {
	if c.config.CacheDir == "" {
		return
	}

	cacheFile := filepath.Join(c.config.CacheDir, "update_check.json")

	// Ensure cache directory exists
	if err := os.MkdirAll(c.config.CacheDir, 0755); err != nil {
		return
	}

	entry := cacheEntry{
		Data:      info,
		CreatedAt: time.Now(),
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return
	}

	_ = os.WriteFile(cacheFile, data, 0644)
}

// getCached retrieves cached update info if still valid
func (c *Checker) getCached() (*Info, bool) {
	if c.config.CacheDir == "" {
		return nil, false
	}

	cacheFile := filepath.Join(c.config.CacheDir, "update_check.json")

	data, err := os.ReadFile(cacheFile)
	if err != nil {
		return nil, false
	}

	var entry cacheEntry
	if err := json.Unmarshal(data, &entry); err != nil {
		return nil, false
	}

	// Check if expired
	if time.Since(entry.CreatedAt) > c.config.CacheTTL {
		_ = os.Remove(cacheFile)
		return nil, false
	}

	return entry.Data, true
}

// defaultCacheDir returns the default cache directory path
func defaultCacheDir(binaryName string) string {
	// Check environment variable first
	if dir := os.Getenv("CACHE_DIR"); dir != "" {
		return filepath.Join(dir, binaryName)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return filepath.Join(".cache", binaryName)
	}

	return filepath.Join(homeDir, ".cache", binaryName)
}

// isCIEnvironment checks if we're running in a CI environment
func isCIEnvironment() bool {
	ciVars := []string{"CI", "GITHUB_ACTIONS", "GITLAB_CI", "CIRCLECI", "TRAVIS", "JENKINS_URL", "BUILDKITE"}
	for _, v := range ciVars {
		if _, ok := os.LookupEnv(v); ok {
			return true
		}
	}
	return false
}

// normalizeVersion ensures version starts with 'v'
func normalizeVersion(v string) string {
	v = strings.TrimSpace(v)
	if v == "" {
		return "v0.0.0"
	}
	if !strings.HasPrefix(v, "v") && !strings.HasPrefix(v, "V") {
		return "v" + v
	}
	return strings.ToLower(v)
}

// compareVersions compares two semantic versions
func compareVersions(v1, v2 string) int {
	// Remove 'v' prefix
	v1 = strings.TrimPrefix(v1, "v")
	v2 = strings.TrimPrefix(v2, "v")

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

	// Check pre-release status
	if strings.Contains(v1, "-") && !strings.Contains(v2, "-") {
		return -1
	}
	if !strings.Contains(v1, "-") && strings.Contains(v2, "-") {
		return 1
	}

	return 0
}
