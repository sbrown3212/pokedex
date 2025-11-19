package pokecache

import "time"

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	for {
		<-ticker.C

		currentTime := time.Now().UTC()

		var entries []struct {
			key       string
			createdAt time.Time
		}

		// Get all entriees in cache
		c.mu.Lock()
		for key, entry := range c.entries {
			entries = append(entries, struct {
				key       string
				createdAt time.Time
			}{
				key:       key,
				createdAt: entry.createdAt,
			})
		}
		c.mu.Unlock()

		// Get keys for entries to be deleted
		var deleteKeys []string
		for _, entry := range entries {
			expired := currentTime.Sub(entry.createdAt) >= c.interval
			if expired {
				deleteKeys = append(deleteKeys, entry.key)
			}
		}

		// Delete expired entries
		c.mu.Lock()
		for _, key := range deleteKeys {
			// Ensure entrie has not been updated
			stillExpired := currentTime.Sub(c.entries[key].createdAt) >= c.interval
			if !stillExpired {
				continue
			}
			delete(c.entries, key)
		}
		c.mu.Unlock()
	}
}
