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
	current, ok := s.items[key]
	if ok {
		if next.Version < current.Version {
			return false
		}
		if next.Version == current.Version && current.Status == "complete" && next.Status != "complete" {
			return false
		}
	}
	s.items[key] = next
	return true
}

func (s *Store) Get(key string) State {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.items[key]
}
