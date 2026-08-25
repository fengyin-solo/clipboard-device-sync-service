package service

import (
	"sort"
	"time"

	"clipboard/internal/model"
	"clipboard/internal/store"
	"clipboard/pkg/idgen"
)

func (s *Service) CreateSyncLog(input model.SyncLog) (*model.SyncLog, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	if _, err := s.store.GetDevice(input.DeviceID); err != nil {
		if err == store.ErrNotFound {
			return nil, model.NewValidationError("device_id", "设备不存在")
		}
		return nil, err
	}
	now := time.Now()
	input.ID = idgen.Hex()
	input.SyncedAt = now
	if err := s.store.CreateSyncLog(&input); err != nil {
		return nil, err
	}
	return &input, nil
}

func (s *Service) GetSyncLog(id string) (*model.SyncLog, error) {
	return s.store.GetSyncLog(id)
}

func (s *Service) ListSyncLogs(filter model.SyncLogFilter, page, size int) ([]*model.SyncLog, int, error) {
	all := s.store.ListSyncLogs()
	matched := make([]*model.SyncLog, 0, len(all))
	for _, sl := range all {
		if filter.Match(sl) {
			matched = append(matched, sl)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].SyncedAt.After(matched[j].SyncedAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.SyncLog{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

func (s *Service) UpdateSyncLog(id string, input model.SyncLog) (*model.SyncLog, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	old, err := s.store.GetSyncLog(id)
	if err != nil {
		return nil, err
	}
	old.DeviceID = input.DeviceID
	old.Direction = input.Direction
	old.EntryCount = input.EntryCount
	old.Status = input.Status
	old.Message = input.Message
	if err := s.store.UpdateSyncLog(old); err != nil {
		return nil, err
	}
	return old, nil
}

func (s *Service) DeleteSyncLog(id string) error {
	return s.store.DeleteSyncLog(id)
}
