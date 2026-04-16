package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu    *sync.Mutex
	cache map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.cache[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.cache[key]
	if exists == false {
		return []byte{}, exists
	}

	return entry.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mu.Lock()

		for key, cacheEntry := range c.cache {
			if time.Since(cacheEntry.createdAt) > interval {
				delete(c.cache, key)
			}
		}

		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) Cache {
	pokecache := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}

	go pokecache.reapLoop(interval)

	return pokecache
}
