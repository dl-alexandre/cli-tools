package cache

import (
	"testing"
	"time"
)

func TestNewMemory(t *testing.T) {
	c := NewMemory(MemoryOptions{
		MaxSize:    100,
		DefaultTTL: 1 * time.Hour,
	})

	if c == nil {
		t.Fatal("NewMemory() returned nil")
	}
}

func TestMemoryCacheSetAndGet(t *testing.T) {
	c := NewMemory(MemoryOptions{
		MaxSize:    100,
		DefaultTTL: 1 * time.Hour,
	})

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

func TestMemoryCacheExpiration(t *testing.T) {
	c := NewMemory(MemoryOptions{
		MaxSize:    100,
		DefaultTTL: 1 * time.Hour,
	})

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

func TestMemoryCacheDelete(t *testing.T) {
	c := NewMemory(MemoryOptions{
		MaxSize:    100,
		DefaultTTL: 1 * time.Hour,
	})

	// Set value
	_ = c.Set("key1", "value1", 0)

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

func TestMemoryCacheClear(t *testing.T) {
	c := NewMemory(MemoryOptions{
		MaxSize:    100,
		DefaultTTL: 1 * time.Hour,
	})

	// Set multiple values
	_ = c.Set("key1", "value1", 0)
	_ = c.Set("key2", "value2", 0)
	_ = c.Set("key3", "value3", 0)

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

func TestMemoryCacheKeys(t *testing.T) {
	c := NewMemory(MemoryOptions{
		MaxSize:    100,
		DefaultTTL: 1 * time.Hour,
	})

	// Set multiple values
	_ = c.Set("key1", "value1", 0)
	_ = c.Set("key2", "value2", 0)
	_ = c.Set("key3", "value3", 0)

	// Get keys
	keys := c.Keys()
	if len(keys) != 3 {
		t.Errorf("Keys() returned %d keys, want 3", len(keys))
	}
}

func TestMemoryCacheLRUEviction(t *testing.T) {
	c := NewMemory(MemoryOptions{
		MaxSize:    3, // Small cache
		DefaultTTL: 1 * time.Hour,
	})

	// Fill cache to capacity
	_ = c.Set("key1", "value1", 0)
	_ = c.Set("key2", "value2", 0)
	_ = c.Set("key3", "value3", 0)

	// Add one more - should evict least recently used (key1)
	_ = c.Set("key4", "value4", 0)

	// key1 should be evicted
	_, ok := c.Get("key1")
	if ok {
		t.Error("Get(key1) should return false after LRU eviction")
	}

	// key2, key3, key4 should still exist
	for _, key := range []string{"key2", "key3", "key4"} {
		_, ok := c.Get(key)
		if !ok {
			t.Errorf("Get(%s) should return true", key)
		}
	}
}

func TestMemoryCacheUpdateExisting(t *testing.T) {
	c := NewMemory(MemoryOptions{
		MaxSize:    100,
		DefaultTTL: 1 * time.Hour,
	})

	// Set initial value
	_ = c.Set("key1", "value1", 0)

	// Update it
	err := c.Set("key1", "value2", 0)
	if err != nil {
		t.Errorf("Set() error updating existing key = %v", err)
	}

	// Verify updated
	value, ok := c.Get("key1")
	if !ok {
		t.Fatal("Get() returned false for updated key")
	}
	if value != "value2" {
		t.Errorf("Get() value = %v, want %v", value, "value2")
	}
}

func TestMemoryCacheStats(t *testing.T) {
	c := NewMemory(MemoryOptions{
		MaxSize:    100,
		DefaultTTL: 1 * time.Hour,
	})

	// Add some items
	_ = c.Set("key1", "value1", 0)
	_ = c.Set("key2", "value2", 0)
	_ = c.Set("key3", "value3", 0)

	mc := c.(*MemoryCache)
	stats := mc.Stats()

	if stats.Size != 3 {
		t.Errorf("Stats.Size = %d, want 3", stats.Size)
	}
	if stats.MaxSize != 100 {
		t.Errorf("Stats.MaxSize = %d, want 100", stats.MaxSize)
	}

	// Test String()
	str := stats.String()
	if str != "Cache: 3/100 items" {
		t.Errorf("Stats.String() = %s, want 'Cache: 3/100 items'", str)
	}
}

func TestMemoryCacheUnlimitedSize(t *testing.T) {
	c := NewMemory(MemoryOptions{
		MaxSize:    0, // Unlimited
		DefaultTTL: 1 * time.Hour,
	})

	// Add many items
	for i := 0; i < 1000; i++ {
		err := c.Set(string(rune(i)), i, 0)
		if err != nil {
			t.Errorf("Set() error with unlimited cache = %v", err)
		}
	}

	// Verify all exist
	keys := c.Keys()
	if len(keys) != 1000 {
		t.Errorf("Keys() = %d, want 1000 (unlimited cache)", len(keys))
	}

	// Check stats string
	mc := c.(*MemoryCache)
	stats := mc.Stats()
	str := stats.String()
	if str != "Cache: 1000 items" {
		t.Errorf("Stats.String() with unlimited = %s, want 'Cache: 1000 items'", str)
	}
}

func TestMemoryCacheComplexValues(t *testing.T) {
	c := NewMemory(MemoryOptions{
		MaxSize:    100,
		DefaultTTL: 1 * time.Hour,
	})

	// Test with struct
	type testStruct struct {
		Name  string
		Count int
	}

	data := testStruct{Name: "test", Count: 42}
	err := c.Set("struct", data, 0)
	if err != nil {
		t.Errorf("Set() error with struct = %v", err)
	}

	value, ok := c.Get("struct")
	if !ok {
		t.Fatal("Get() returned false for struct")
	}

	// Memory cache preserves types
	result, ok := value.(testStruct)
	if !ok {
		t.Fatalf("Get() returned wrong type: %T", value)
	}
	if result.Name != "test" || result.Count != 42 {
		t.Errorf("Get() returned wrong values: %+v", result)
	}
}

func TestMemoryCacheConcurrentAccess(t *testing.T) {
	c := NewMemory(MemoryOptions{
		MaxSize:    1000,
		DefaultTTL: 1 * time.Hour,
	})

	// This test verifies the mutex doesn't cause panics
	// Run operations in goroutines
	done := make(chan bool, 10)

	// Writer
	go func() {
		for i := 0; i < 100; i++ {
			_ = c.Set(string(rune(i)), i, 0)
		}
		done <- true
	}()

	// Reader
	go func() {
		for i := 0; i < 100; i++ {
			c.Get(string(rune(i)))
		}
		done <- true
	}()

	// Wait for completion
	<-done
	<-done

	// If we get here without panic, mutex is working
	t.Log("Concurrent access completed without panic")
}
