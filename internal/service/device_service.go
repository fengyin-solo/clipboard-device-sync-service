package service

import (
	"sort"
	"time"

	"clipboard/internal/model"
	"clipboard/pkg/idgen"
)

func (s *Service) CreateDevice(input model.Device) (*model.Device, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	now := time.Now()
	input.ID = idgen.Hex()
	input.CreatedAt = now
	input.LastSyncAt = now
	if err := s.store.CreateDevice(&input); err != nil {
		return nil, err
	}
	return &input, nil
}

func (s *Service) GetDevice(id string) (*model.Device, error) {
	return s.store.GetDevice(id)
}

func (s *Service) ListDevices(filter model.DeviceFilter, page, size int) ([]*model.Device, int, error) {
	all := s.store.ListDevices()
	matched := make([]*model.Device, 0, len(all))
	for _, d := range all {
		if filter.Match(d) {
			matched = append(matched, d)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].CreatedAt.After(matched[j].CreatedAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.Device{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

func (s *Service) UpdateDevice(id string, input model.Device) (*model.Device, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	old, err := s.store.GetDevice(id)
	if err != nil {
		return nil, err
	}
	old.Name = input.Name
	old.Platform = input.Platform
	old.Token = input.Token
	if err := s.store.UpdateDevice(old); err != nil {
		return nil, err
	}
	return old, nil
}

func (s *Service) DeleteDevice(id string) error {
	return s.store.DeleteDevice(id)
}
