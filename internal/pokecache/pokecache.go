package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	interval time.Duration
	mu       *sync.Mutex
}

type cacheEntry struct {
	createAt time.Time
	val      []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cacheMap: make(map[string]cacheEntry),
		interval: interval,
		mu:       &sync.Mutex{},
	}
	go cache.readLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		createAt: time.Now(),
		val:      val,
	}
	c.cacheMap[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if cacheEntry, ok := c.cacheMap[key]; ok {
		return cacheEntry.val, true
	}

	return nil, false
}

func (c *Cache) readLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.cacheMap {
		if v.createAt.Before(now.Add(-last)) {
			delete(c.cacheMap, k)
		}
	}
}
