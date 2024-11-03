package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	Entries  map[string]*cacheEntry
	Mu       sync.Mutex
	Interval time.Duration
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	c.Entries[key] = &cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	if entry, ok := c.Entries[key]; ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) ReapLoop() {
	ticker := time.NewTicker(c.Interval)
	defer ticker.Stop()
	for range ticker.C {
		c.Mu.Lock()
		for key, entry := range c.Entries {
			if time.Since(entry.createdAt) > c.Interval {
				delete(c.Entries, key)
			}
		}
		c.Mu.Unlock()
	}
}

func NewCache(interval time.Duration) Cache {
	return Cache{
		Entries:  make(map[string]*cacheEntry),
		Interval: interval,
	}
}
