package store

import (
	"clipboard/internal/model"
)

func (s *MemoryStore) CreateGroup(g *model.Group) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, exist := range s.groups {
		if exist.Name == g.Name {
			return ErrConflict
		}
	}
	s.groups[g.ID] = g
	return nil
}

func (s *MemoryStore) GetGroup(id string) (*model.Group, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	g, ok := s.groups[id]
	if !ok {
		return nil, ErrNotFound
	}
	return g, nil
}

func (s *MemoryStore) ListGroups() []*model.Group {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.Group, 0, len(s.groups))
	for _, g := range s.groups {
		list = append(list, g)
	}
	return list
}

func (s *MemoryStore) UpdateGroup(g *model.Group) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.groups[g.ID]; !ok {
		return ErrNotFound
	}
	for _, exist := range s.groups {
		if exist.ID != g.ID && exist.Name == g.Name {
			return ErrConflict
		}
	}
	s.groups[g.ID] = g
	return nil
}

func (s *MemoryStore) DeleteGroup(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.groups[id]; !ok {
		return ErrNotFound
	}
	delete(s.groups, id)
	return nil
}
