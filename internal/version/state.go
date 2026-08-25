package version

import "sync"

type State struct {
	Version int
	Status  string
}

type Store struct {
	mu    sync.RWMutex
	items map[string]State
}

func NewStore() *Store { return &Store{items: make(map[string]State)} }

func (s *Store) Put(key string, next State) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[key] = next
	return true
}

func (s *Store) Get(key string) State {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.items[key]
}
