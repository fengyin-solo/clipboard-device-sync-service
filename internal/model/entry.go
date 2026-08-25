package model

import (
	"strings"
	"time"
)

const (
	EntryContentTypeText  = "text"
	EntryContentTypeImage = "image"
	EntryContentTypeLink  = "link"
	EntryContentTypeCode  = "code"
)

const (
	EntryStatusPending  = "pending"
	EntryStatusSynced   = "synced"
	EntryStatusConflict = "conflict"
)

var entryTransitions = map[string]map[string]bool{
	EntryStatusPending: {EntryStatusSynced: true},
	EntryStatusSynced:  {EntryStatusConflict: true},
}

func EntryCanTransition(from, to string) bool {
	if m, ok := entryTransitions[from]; ok {
		return m[to]
	}
	return false
}

type Entry struct {
	ID           string    `json:"id"`
	Content      string    `json:"content"`
	ContentType  string    `json:"content_type"`
	GroupID      string    `json:"group_id"`
	SizeBytes    int64     `json:"size_bytes"`
	Pinned       bool      `json:"pinned"`
	SourceDevice string    `json:"source_device"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ExpireAt     time.Time `json:"expire_at"`
}

func (e *Entry) Validate() error {
	e.Content = strings.TrimSpace(e.Content)
	if e.Content == "" {
		return NewValidationError("content", "内容不能为空")
	}
	if e.ContentType == "" {
		e.ContentType = EntryContentTypeText
	}
	if e.ContentType != EntryContentTypeText && e.ContentType != EntryContentTypeImage &&
		e.ContentType != EntryContentTypeLink && e.ContentType != EntryContentTypeCode {
		return NewValidationError("content_type", "内容类型不合法")
	}
	if e.Status == "" {
		e.Status = EntryStatusPending
	}
	if e.Status != EntryStatusPending && e.Status != EntryStatusSynced && e.Status != EntryStatusConflict {
		return NewValidationError("status", "状态不合法")
	}
	if e.SizeBytes < 0 {
		return NewValidationError("size_bytes", "大小不能为负数")
	}
	return nil
}

type EntryFilter struct {
	GroupID     string
	ContentType string
	Status      string
	Keyword     string
	Pinned      *bool
}

func (f EntryFilter) Match(e *Entry) bool {
	if f.GroupID != "" && e.GroupID != f.GroupID {
		return false
	}
	if f.ContentType != "" && e.ContentType != f.ContentType {
		return false
	}
	if f.Status != "" && e.Status != f.Status {
		return false
	}
	if f.Pinned != nil && e.Pinned != *f.Pinned {
		return false
	}
	if f.Keyword != "" {
		k := strings.ToLower(strings.TrimSpace(f.Keyword))
		if k != "" && !strings.Contains(strings.ToLower(e.Content), k) {
			return false
		}
	}
	return true
}
