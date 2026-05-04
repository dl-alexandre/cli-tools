// Package cache provides standardized caching for CLI applications.
// Supports both file-based and in-memory caching with consistent interfaces.
//
// Example usage:
//   // File-based cache (default)
//   c := cache.New(cache.DefaultDir("myapp"), 24*time.Hour)
//   c.Set("key", data, 1*time.Hour)
//
//   // Memory cache for hot data
//   mc := cache.NewMemory(cache.MemoryOptions{MaxSize: 100})
//   mc.Set("key", data, 5*time.Minute)
package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"crypto/sha256"
)

// Cache defines the interface for caching operations
type Cache interface {
	// Get retrieves an item from the cache
	// Returns the cached value and true if found and not expired
	Get(key string) (interface{}, bool)

	// Set stores an item in the cache with the specified TTL
	Set(key string, value interface{}, ttl time.Duration) error

	// Delete removes an item from the cache
	Delete(key string) error

	// Clear removes all items from the cache
	Clear() error

	// Keys returns all keys in the cache
	Keys() []string
}

// FileCache implements file-based caching
type FileCache struct {
	dir string
	defaultTTL time.Duration
}

// New creates a new file-based cache
// dir: cache directory path
// defaultTTL: default time-to-live for cached items
func New(dir string, defaultTTL time.Duration) Cache {
	return &FileCache{
		dir: dir,
		defaultTTL: defaultTTL,
	}
}

// cacheEntry represents a cached item
type cacheEntry struct {
	Data      json.RawMessage `json:"data"`
	CreatedAt time.Time       `json:"created_at"`
	TTL       time.Duration   `json:"ttl"`
}

// Get retrieves an item from the cache
func (c *FileCache) Get(key string) (interface{}, bool) {
	path := c.filePath(key)

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, false
	}

	var entry cacheEntry
	if err := json.Unmarshal(data, &entry); err != nil {
		return nil, false
	}

	// Check if expired
	if time.Since(entry.CreatedAt) > entry.TTL {
		_ = os.Remove(path)
		return nil, false
	}

	var value interface{}
	if err := json.Unmarshal(entry.Data, &value); err != nil {
		return nil, false
	}

	return value, true
}

// Set stores an item in the cache
func (c *FileCache) Set(key string, value interface{}, ttl time.Duration) error {
	if ttl == 0 {
		ttl = c.defaultTTL
	}

	path := c.filePath(key)

	// Ensure cache directory exists
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create cache directory: %w", err)
	}

	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to encode cache value: %w", err)
	}

	entry := cacheEntry{
		Data:      data,
		CreatedAt: time.Now(),
		TTL:       ttl,
	}

	encoded, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("failed to encode cache entry: %w", err)
	}

	if err := os.WriteFile(path, encoded, 0644); err != nil {
		return fmt.Errorf("failed to write cache file: %w", err)
	}

	return nil
}

// Delete removes an item from the cache
func (c *FileCache) Delete(key string) error {
	path := c.filePath(key)
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete cache file: %w", err)
	}
	return nil
}

// Clear removes all items from the cache
func (c *FileCache) Clear() error {
	entries, err := os.ReadDir(c.dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to read cache directory: %w", err)
	}

	for _, entry := range entries {
		if err := os.Remove(filepath.Join(c.dir, entry.Name())); err != nil {
			return fmt.Errorf("failed to remove cache file: %w", err)
		}
	}

	return nil
}

// Keys returns all keys in the cache
func (c *FileCache) Keys() []string {
	entries, err := os.ReadDir(c.dir)
	if err != nil {
		return []string{}
	}

	keys := make([]string, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			// Remove .json extension if present
			name := entry.Name()
			if filepath.Ext(name) == ".json" {
				name = name[:len(name)-5]
			}
			keys = append(keys, name)
		}
	}

	return keys
}

// filePath generates the file path for a cache key
func (c *FileCache) filePath(key string) string {
	// Sanitize key for use as filename
	safeKey := sanitizeKey(key)
	return filepath.Join(c.dir, safeKey+".json")
}

// sanitizeKey makes a key safe for use as a filename
func sanitizeKey(key string) string {
	// Use SHA256 hash for long or complex keys
	if len(key) > 50 {
		hash := sha256.Sum256([]byte(key))
		return fmt.Sprintf("%x", hash[:16])
	}

	// Replace unsafe characters
	result := make([]byte, 0, len(key))
	for _, c := range key {
		switch {
		case (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9'):
			result = append(result, byte(c))
		case c == ' ' || c == '/':
			result = append(result, '_')
		default:
			result = append(result, '-')
		}
	}

	return string(result)
}

// DefaultDir returns the default cache directory for an application
// Uses XDG Base Directory specification: ~/.cache/<appname>
func DefaultDir(appName string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		return filepath.Join(".cache", appName)
	}

	// Check XDG_CACHE_HOME
	if cacheHome := os.Getenv("XDG_CACHE_HOME"); cacheHome != "" {
		return filepath.Join(cacheHome, appName)
	}

	return filepath.Join(home, ".cache", appName)
}

// GenerateKey creates a cache key from a prefix and parameters
// Useful for API response caching
func GenerateKey(prefix string, params map[string]interface{}) string {
	if len(params) == 0 {
		return prefix
	}

	// Sort params for consistent keys
	paramBytes, _ := json.Marshal(params)
	hash := sha256.Sum256(paramBytes)
	return fmt.Sprintf("%s_%x", prefix, hash[:8])
}
