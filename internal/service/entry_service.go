package service

import (
	"sort"
	"time"

	"clipboard/internal/model"
	"clipboard/internal/store"
	"clipboard/pkg/idgen"
)

func (s *Service) CreateEntry(input model.Entry) (*model.Entry, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	if input.GroupID != "" {
		if _, err := s.store.GetGroup(input.GroupID); err != nil {
			if err == store.ErrNotFound {
				return nil, model.NewValidationError("group_id", "分组不存在")
			}
			return nil, err
		}
	}
	now := time.Now()
	input.ID = idgen.Hex()
	input.CreatedAt = now
	input.UpdatedAt = now
	if input.ExpireAt.IsZero() {
		input.ExpireAt = now.Add(24 * time.Hour * 7)
	}
	if err := s.store.CreateEntry(&input); err != nil {
		return nil, err
	}
	return &input, nil
}

func (s *Service) GetEntry(id string) (*model.Entry, error) {
	return s.store.GetEntry(id)
}

func (s *Service) ListEntries(filter model.EntryFilter, page, size int) ([]*model.Entry, int, error) {
	all := s.store.ListEntries()
	matched := make([]*model.Entry, 0, len(all))
	for _, e := range all {
		if filter.Match(e) {
			matched = append(matched, e)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		if matched[i].Pinned != matched[j].Pinned {
			return matched[i].Pinned
		}
		return matched[i].CreatedAt.After(matched[j].CreatedAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.Entry{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

func (s *Service) UpdateEntry(id string, input model.Entry) (*model.Entry, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	old, err := s.store.GetEntry(id)
	if err != nil {
		return nil, err
	}
	if input.Status != old.Status {
		if !model.EntryCanTransition(old.Status, input.Status) {
			return nil, model.NewValidationError("status", "非法的状态流转")
		}
	}
	if input.GroupID != "" && input.GroupID != old.GroupID {
		if _, err := s.store.GetGroup(input.GroupID); err != nil {
			if err == store.ErrNotFound {
				return nil, model.NewValidationError("group_id", "分组不存在")
			}
			return nil, err
		}
	}
	old.Content = input.Content
	old.ContentType = input.ContentType
	old.GroupID = input.GroupID
	old.SizeBytes = input.SizeBytes
	old.Pinned = input.Pinned
	old.SourceDevice = input.SourceDevice
	old.Status = input.Status
	old.UpdatedAt = time.Now()
	if !input.ExpireAt.IsZero() {
		old.ExpireAt = input.ExpireAt
	}
	if err := s.store.UpdateEntry(old); err != nil {
		return nil, err
	}
	return old, nil
}

func (s *Service) DeleteEntry(id string) error {
	return s.store.DeleteEntry(id)
}

func (s *Service) BatchCreateEntries(inputs []model.Entry) ([]*model.Entry, error) {
	results := make([]*model.Entry, 0, len(inputs))
	for _, input := range inputs {
		out, err := s.CreateEntry(input)
		if err != nil {
			return nil, err
		}
		results = append(results, out)
	}
	return results, nil
}

func (s *Service) BatchDeleteEntries(ids []string) error {
	for _, id := range ids {
		if err := s.store.DeleteEntry(id); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) BatchUpdateEntryStatus(ids []string, status string) error {
	for _, id := range ids {
		old, err := s.store.GetEntry(id)
		if err != nil {
			return err
		}
		if !model.EntryCanTransition(old.Status, status) {
			return model.NewValidationError("status", "非法的状态流转: "+old.Status+" -> "+status)
		}
		old.Status = status
		old.UpdatedAt = time.Now()
		if err := s.store.UpdateEntry(old); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) BatchPinEntries(ids []string, pinned bool) error {
	for _, id := range ids {
		old, err := s.store.GetEntry(id)
		if err != nil {
			return err
		}
		old.Pinned = pinned
		old.UpdatedAt = time.Now()
		if err := s.store.UpdateEntry(old); err != nil {
			return err
		}
	}
	return nil
}
