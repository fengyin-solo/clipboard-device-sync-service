package service

import (
	"time"

	"clipboard/internal/model"
	"clipboard/internal/store"
	"clipboard/pkg/idgen"
)

func (s *Service) CreateEntryTag(input model.EntryTag) (*model.EntryTag, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	if _, err := s.store.GetEntry(input.EntryID); err != nil {
		if err == store.ErrNotFound {
			return nil, model.NewValidationError("entry_id", "条目不存在")
		}
		return nil, err
	}
	if _, err := s.store.GetTag(input.TagID); err != nil {
		if err == store.ErrNotFound {
			return nil, model.NewValidationError("tag_id", "标签不存在")
		}
		return nil, err
	}
	now := time.Now()
	input.ID = idgen.Hex()
	input.CreatedAt = now
	if err := s.store.CreateEntryTag(&input); err != nil {
		return nil, err
	}
	return &input, nil
}

func (s *Service) GetEntryTag(id string) (*model.EntryTag, error) {
	return s.store.GetEntryTag(id)
}

func (s *Service) ListEntryTags(filter model.EntryTagFilter) ([]*model.EntryTag, error) {
	all := s.store.ListEntryTags()
	matched := make([]*model.EntryTag, 0, len(all))
	for _, et := range all {
		if filter.Match(et) {
			matched = append(matched, et)
		}
	}
	return matched, nil
}

func (s *Service) DeleteEntryTag(id string) error {
	return s.store.DeleteEntryTag(id)
}
