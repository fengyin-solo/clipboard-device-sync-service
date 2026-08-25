package document

import "sync"

type Cache struct {
	mu    sync.RWMutex
	items map[string]*Document
}

func NewCache() *Cache { return &Cache{items: make(map[string]*Document)} }

func (c *Cache) Put(document *Document) {
	if document == nil {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	copyDocument := *document
	c.items[document.ID] = &copyDocument
}

func (c *Cache) Get(id string) (*Document, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	document, ok := c.items[id]
	if !ok {
		return nil, false
	}
	copyDocument := *document
	return &copyDocument, true
}
