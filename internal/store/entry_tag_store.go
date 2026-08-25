package store

import (
	"clipboard/internal/model"
)

func (s *MemoryStore) CreateEntryTag(et *model.EntryTag) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, exist := range s.entryTags {
		if exist.EntryID == et.EntryID && exist.TagID == et.TagID {
			return ErrConflict
		}
	}
	s.entryTags[et.ID] = et
	return nil
}

func (s *MemoryStore) GetEntryTag(id string) (*model.EntryTag, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	et, ok := s.entryTags[id]
	if !ok {
		return nil, ErrNotFound
	}
	return et, nil
}

func (s *MemoryStore) ListEntryTags() []*model.EntryTag {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.EntryTag, 0, len(s.entryTags))
	for _, et := range s.entryTags {
		list = append(list, et)
	}
	return list
}

func (s *MemoryStore) DeleteEntryTag(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.entryTags[id]; !ok {
		return ErrNotFound
	}
	delete(s.entryTags, id)
	return nil
}

func (s *MemoryStore) ListEntryTagsByEntry(entryID string) []*model.EntryTag {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.EntryTag, 0)
	for _, et := range s.entryTags {
		if et.EntryID == entryID {
			list = append(list, et)
		}
	}
	return list
}

func (s *MemoryStore) ListEntryTagsByTag(tagID string) []*model.EntryTag {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.EntryTag, 0)
	for _, et := range s.entryTags {
		if et.TagID == tagID {
			list = append(list, et)
		}
	}
	return list
}
