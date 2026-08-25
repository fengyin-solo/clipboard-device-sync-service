package store

import (
	"clipboard/internal/model"
)

func (s *MemoryStore) CreateEntry(e *model.Entry) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, exist := range s.entries {
		if exist.Content == e.Content && exist.SourceDevice == e.SourceDevice && exist.GroupID == e.GroupID {
			return ErrConflict
		}
	}
	s.entries[e.ID] = e
	return nil
}

func (s *MemoryStore) GetEntry(id string) (*model.Entry, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	e, ok := s.entries[id]
	if !ok {
		return nil, ErrNotFound
	}
	return e, nil
}

func (s *MemoryStore) ListEntries() []*model.Entry {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.Entry, 0, len(s.entries))
	for _, e := range s.entries {
		list = append(list, e)
	}
	return list
}

func (s *MemoryStore) UpdateEntry(e *model.Entry) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.entries[e.ID]; !ok {
		return ErrNotFound
	}
	for _, exist := range s.entries {
		if exist.ID != e.ID && exist.Content == e.Content && exist.SourceDevice == e.SourceDevice && exist.GroupID == e.GroupID {
			return ErrConflict
		}
	}
	s.entries[e.ID] = e
	return nil
}

func (s *MemoryStore) DeleteEntry(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.entries[id]; !ok {
		return ErrNotFound
	}
	delete(s.entries, id)
	return nil
}
