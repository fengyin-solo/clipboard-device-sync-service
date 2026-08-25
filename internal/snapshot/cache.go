package snapshot

import "sync"

type Cache struct {
	mu   sync.RWMutex
	data map[string][]byte
}

func NewCache() *Cache {
	return &Cache{data: make(map[string][]byte)}
}

func (c *Cache) Put(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = append([]byte(nil), value...)
}

func (c *Cache) Get(key string) []byte {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return append([]byte(nil), c.data[key]...)
}
