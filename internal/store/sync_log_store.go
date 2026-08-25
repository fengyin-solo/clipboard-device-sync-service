package store

import (
	"clipboard/internal/model"
)

func (s *MemoryStore) CreateSyncLog(sl *model.SyncLog) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncLogs[sl.ID] = sl
	return nil
}

func (s *MemoryStore) GetSyncLog(id string) (*model.SyncLog, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	sl, ok := s.syncLogs[id]
	if !ok {
		return nil, ErrNotFound
	}
	return sl, nil
}

func (s *MemoryStore) ListSyncLogs() []*model.SyncLog {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.SyncLog, 0, len(s.syncLogs))
	for _, sl := range s.syncLogs {
		list = append(list, sl)
	}
	return list
}

func (s *MemoryStore) UpdateSyncLog(sl *model.SyncLog) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.syncLogs[sl.ID]; !ok {
		return ErrNotFound
	}
	s.syncLogs[sl.ID] = sl
	return nil
}

func (s *MemoryStore) DeleteSyncLog(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.syncLogs[id]; !ok {
		return ErrNotFound
	}
	delete(s.syncLogs, id)
	return nil
}
