package cache

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

// MemoryCache implements in-memory caching with LRU eviction
type MemoryCache struct {
	maxSize int
	ttl     time.Duration
	items   map[string]*list.Element
	order   *list.List
	mu      sync.RWMutex
}

// MemoryOptions configures memory cache behavior
type MemoryOptions struct {
	// MaxSize is the maximum number of items to store (0 = unlimited)
	MaxSize int

	// DefaultTTL is the default time-to-live for cached items
	DefaultTTL time.Duration
}

// memoryCacheEntry represents an item in the memory cache
type memoryCacheEntry struct {
	key       string
	value     interface{}
	createdAt time.Time
	ttl       time.Duration
}

// NewMemory creates a new memory-based cache
func NewMemory(opts MemoryOptions) Cache {
	if opts.DefaultTTL == 0 {
		opts.DefaultTTL = 1 * time.Hour
	}

	return &MemoryCache{
		maxSize: opts.MaxSize,
		ttl:     opts.DefaultTTL,
		items:   make(map[string]*list.Element),
		order:   list.New(),
	}
}

// Get retrieves an item from the memory cache
func (c *MemoryCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	elem, ok := c.items[key]
	c.mu.RUnlock()

	if !ok {
		return nil, false
	}

	entry := elem.Value.(*memoryCacheEntry)

	// Check if expired
	if time.Since(entry.createdAt) > entry.ttl {
		c.Delete(key)
		return nil, false
	}

	// Move to front (most recently used)
	c.mu.Lock()
	c.order.MoveToFront(elem)
	c.mu.Unlock()

	return entry.value, true
}

// Set stores an item in the memory cache
func (c *MemoryCache) Set(key string, value interface{}, ttl time.Duration) error {
	if ttl == 0 {
		ttl = c.ttl
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	// Check if key already exists
	if elem, ok := c.items[key]; ok {
		// Update existing entry
		c.order.MoveToFront(elem)
		elem.Value.(*memoryCacheEntry).value = value
		elem.Value.(*memoryCacheEntry).createdAt = time.Now()
		elem.Value.(*memoryCacheEntry).ttl = ttl
		return nil
	}

	// Check if we need to evict
	if c.maxSize > 0 && len(c.items) >= c.maxSize {
		c.evictLRU()
	}

	// Add new entry
	entry := &memoryCacheEntry{
		key:       key,
		value:     value,
		createdAt: time.Now(),
		ttl:       ttl,
	}

	elem := c.order.PushFront(entry)
	c.items[key] = elem

	return nil
}

// Delete removes an item from the memory cache
func (c *MemoryCache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		c.order.Remove(elem)
		delete(c.items, key)
	}

	return nil
}

// Clear removes all items from the memory cache
func (c *MemoryCache) Clear() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[string]*list.Element)
	c.order = list.New()

	return nil
}

// Keys returns all keys in the memory cache
func (c *MemoryCache) Keys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	keys := make([]string, 0, len(c.items))
	for key := range c.items {
		keys = append(keys, key)
	}

	return keys
}

// evictLRU removes the least recently used item
func (c *MemoryCache) evictLRU() {
	elem := c.order.Back()
	if elem == nil {
		return
	}

	entry := elem.Value.(*memoryCacheEntry)
	c.order.Remove(elem)
	delete(c.items, entry.key)
}

// Stats returns cache statistics
func (c *MemoryCache) Stats() MemoryCacheStats {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return MemoryCacheStats{
		Size:    len(c.items),
		MaxSize: c.maxSize,
	}
}

// MemoryCacheStats represents cache statistics
type MemoryCacheStats struct {
	Size    int
	MaxSize int
}

// String returns a string representation of the stats
func (s MemoryCacheStats) String() string {
	if s.MaxSize > 0 {
		return fmt.Sprintf("Cache: %d/%d items", s.Size, s.MaxSize)
	}
	return fmt.Sprintf("Cache: %d items", s.Size)
}
