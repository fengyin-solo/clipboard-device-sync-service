package model

import (
	"time"
)

type EntryTag struct {
	ID        string    `json:"id"`
	EntryID   string    `json:"entry_id"`
	TagID     string    `json:"tag_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (et *EntryTag) Validate() error {
	if et.EntryID == "" {
		return NewValidationError("entry_id", "条目ID不能为空")
	}
	if et.TagID == "" {
		return NewValidationError("tag_id", "标签ID不能为空")
	}
	return nil
}

type EntryTagFilter struct {
	EntryID string
	TagID   string
}

func (f EntryTagFilter) Match(et *EntryTag) bool {
	if f.EntryID != "" && et.EntryID != f.EntryID {
		return false
	}
	if f.TagID != "" && et.TagID != f.TagID {
		return false
	}
	return true
}
