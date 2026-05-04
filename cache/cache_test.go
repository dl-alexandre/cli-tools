package cache

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tempDir := t.TempDir()
	cache := New(tempDir, 1*time.Hour)

	if cache == nil {
		t.Fatal("New() returned nil")
	}
}

func TestFileCacheSetAndGet(t *testing.T) {
	tempDir := t.TempDir()
	c := New(tempDir, 1*time.Hour)

	// Test Set
	err := c.Set("key1", "value1", 0) // Uses default TTL
	if err != nil {
		t.Errorf("Set() error = %v", err)
	}

	// Test Get
	value, ok := c.Get("key1")
	if !ok {
		t.Error("Get() returned false for existing key")
	}
	if value != "value1" {
		t.Errorf("Get() value = %v, want %v", value, "value1")
	}
}

func TestFileCacheExpiration(t *testing.T) {
	tempDir := t.TempDir()
	c := New(tempDir, 1*time.Hour)

	// Set with very short TTL
	err := c.Set("key1", "value1", 1*time.Millisecond)
	if err != nil {
		t.Errorf("Set() error = %v", err)
	}

	// Should exist immediately
	_, ok := c.Get("key1")
	if !ok {
		t.Error("Get() should return true immediately after Set()")
	}

	// Wait for expiration
	time.Sleep(10 * time.Millisecond)

	// Should be expired now
	_, ok = c.Get("key1")
	if ok {
		t.Error("Get() should return false for expired key")
	}
}

func TestFileCacheDelete(t *testing.T) {
	tempDir := t.TempDir()
	c := New(tempDir, 1*time.Hour)

	// Set value
	c.Set("key1", "value1", 0)

	// Delete it
	err := c.Delete("key1")
	if err != nil {
		t.Errorf("Delete() error = %v", err)
	}

	// Should not exist
	_, ok := c.Get("key1")
	if ok {
		t.Error("Get() should return false after Delete()")
	}
}

func TestFileCacheClear(t *testing.T) {
	tempDir := t.TempDir()
	c := New(tempDir, 1*time.Hour)

	// Set multiple values
	c.Set("key1", "value1", 0)
	c.Set("key2", "value2", 0)
	c.Set("key3", "value3", 0)

	// Clear all
	err := c.Clear()
	if err != nil {
		t.Errorf("Clear() error = %v", err)
	}

	// All should be gone
	for _, key := range []string{"key1", "key2", "key3"} {
		_, ok := c.Get(key)
		if ok {
			t.Errorf("Get(%s) should return false after Clear()", key)
		}
	}
}

func TestFileCacheKeys(t *testing.T) {
	tempDir := t.TempDir()
	c := New(tempDir, 1*time.Hour)

	// Set multiple values
	c.Set("key1", "value1", 0)
	c.Set("key2", "value2", 0)
	c.Set("key3", "value3", 0)

	// Get keys
	keys := c.Keys()
	if len(keys) != 3 {
		t.Errorf("Keys() returned %d keys, want 3", len(keys))
	}

	// Check all expected keys are present
	keyMap := make(map[string]bool)
	for _, k := range keys {
		keyMap[k] = true
	}

	for _, expected := range []string{"key1", "key2", "key3"} {
		if !keyMap[expected] {
			t.Errorf("Keys() missing expected key: %s", expected)
		}
	}
}

func TestFileCacheComplexValues(t *testing.T) {
	tempDir := t.TempDir()
	c := New(tempDir, 1*time.Hour)

	// Test with map
	data := map[string]interface{}{
		"name":  "test",
		"count": 42,
		"items": []string{"a", "b", "c"},
	}

	err := c.Set("complex", data, 0)
	if err != nil {
		t.Errorf("Set() error = %v", err)
	}

	value, ok := c.Get("complex")
	if !ok {
		t.Fatal("Get() returned false for complex value")
	}

	// The value will be unmarshaled as interface{}
	// We just verify it's not nil and contains data
	if value == nil {
		t.Error("Get() returned nil for complex value")
	}
}

func TestDefaultDir(t *testing.T) {
	appName := "testapp"
	dir := DefaultDir(appName)

	// Should contain app name
	if !contains(dir, appName) {
		t.Errorf("DefaultDir() should contain app name, got %s", dir)
	}

	// Should contain .cache or follow XDG
	if !contains(dir, ".cache") && os.Getenv("XDG_CACHE_HOME") == "" {
		t.Errorf("DefaultDir() should contain .cache, got %s", dir)
	}
}

func TestDefaultDirWithXDG(t *testing.T) {
	// Save original
	origXDG := os.Getenv("XDG_CACHE_HOME")
	defer os.Setenv("XDG_CACHE_HOME", origXDG)

	// Set XDG_CACHE_HOME
	os.Setenv("XDG_CACHE_HOME", "/custom/cache")

	dir := DefaultDir("testapp")
	expected := "/custom/cache/testapp"

	if dir != expected {
		t.Errorf("DefaultDir() with XDG = %s, want %s", dir, expected)
	}
}

func TestGenerateKey(t *testing.T) {
	tests := []struct {
		prefix string
		params map[string]interface{}
		want   string
	}{
		{
			prefix: "api",
			params: nil,
			want:   "api",
		},
		{
			prefix: "users",
			params: map[string]interface{}{"id": 123},
			want:   "users_",
		},
		{
			prefix: "search",
			params: map[string]interface{}{"q": "test", "limit": 10},
			want:   "search_",
		},
	}

	for _, tt := range tests {
		got := GenerateKey(tt.prefix, tt.params)
		if tt.params == nil {
			if got != tt.want {
				t.Errorf("GenerateKey(%s, nil) = %s, want %s", tt.prefix, got, tt.want)
			}
		} else {
			// With params, should have prefix_ and hash
			if len(got) <= len(tt.prefix)+1 {
				t.Errorf("GenerateKey(%s, %v) = %s, want prefix with hash", tt.prefix, tt.params, got)
			}
		}
	}
}

func TestSanitizeKey(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"simple", "simple"},
		{"with spaces", "with_spaces"},
		{"with/slash", "with_slash"},
		{"UPPER", "UPPER"},
		{"123", "123"},
	}

	for _, tt := range tests {
		result := sanitizeKey(tt.input)
		if len(tt.input) <= 50 && result != tt.expected {
			t.Errorf("sanitizeKey(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestSanitizeKeyLong(t *testing.T) {
	// Long keys should be hashed
	longKey := "this is a very long key that exceeds fifty characters and should be hashed"
	result := sanitizeKey(longKey)

	// Result should be a hex hash (32 chars for first 16 bytes of SHA256)
	if len(result) != 32 {
		t.Errorf("sanitizeKey(long) = %s (len=%d), expected 32 char hash", result, len(result))
	}
}

func TestFileCacheNonExistentKey(t *testing.T) {
	tempDir := t.TempDir()
	c := New(tempDir, 1*time.Hour)

	value, ok := c.Get("non-existent")
	if ok {
		t.Error("Get() should return false for non-existent key")
	}
	if value != nil {
		t.Error("Get() should return nil for non-existent key")
	}
}

func TestFileCacheDirectoryCreation(t *testing.T) {
	tempDir := t.TempDir()
	deepDir := filepath.Join(tempDir, "a", "b", "c")
	c := New(deepDir, 1*time.Hour)

	err := c.Set("key", "value", 0)
	if err != nil {
		t.Errorf("Set() should create directories, error = %v", err)
	}

	// Verify directory was created
	if _, err := os.Stat(deepDir); os.IsNotExist(err) {
		t.Error("Set() did not create cache directories")
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsSubstring(s, substr))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
