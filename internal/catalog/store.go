package catalog

import "sync"

type Store struct {
	mu    sync.RWMutex
	items map[string]string
}

func NewStore() *Store { return &Store{items: make(map[string]string)} }

func (s *Store) Put(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[key] = value
}

func (s *Store) Snapshot() map[string]string {
	return s.items
}
