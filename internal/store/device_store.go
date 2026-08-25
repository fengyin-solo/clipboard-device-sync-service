package store

import (
	"clipboard/internal/model"
)

func (s *MemoryStore) CreateDevice(d *model.Device) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, exist := range s.devices {
		if exist.Token == d.Token && d.Token != "" {
			return ErrConflict
		}
	}
	s.devices[d.ID] = d
	return nil
}

func (s *MemoryStore) GetDevice(id string) (*model.Device, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	d, ok := s.devices[id]
	if !ok {
		return nil, ErrNotFound
	}
	return d, nil
}

func (s *MemoryStore) ListDevices() []*model.Device {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.Device, 0, len(s.devices))
	for _, d := range s.devices {
		list = append(list, d)
	}
	return list
}

func (s *MemoryStore) UpdateDevice(d *model.Device) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.devices[d.ID]; !ok {
		return ErrNotFound
	}
	for _, exist := range s.devices {
		if exist.ID != d.ID && exist.Token == d.Token && d.Token != "" {
			return ErrConflict
		}
	}
	s.devices[d.ID] = d
	return nil
}

func (s *MemoryStore) DeleteDevice(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.devices[id]; !ok {
		return ErrNotFound
	}
	delete(s.devices, id)
	return nil
}
