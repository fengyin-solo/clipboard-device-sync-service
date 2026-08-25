package model

import (
	"strings"
	"time"
)

const (
	SyncLogDirectionPush = "push"
	SyncLogDirectionPull = "pull"
)

const (
	SyncLogStatusSuccess = "success"
	SyncLogStatusFailed  = "failed"
)

type SyncLog struct {
	ID         string    `json:"id"`
	DeviceID   string    `json:"device_id"`
	Direction  string    `json:"direction"`
	EntryCount int       `json:"entry_count"`
	Status     string    `json:"status"`
	Message    string    `json:"message"`
	SyncedAt   time.Time `json:"synced_at"`
}

func (s *SyncLog) Validate() error {
	if s.DeviceID == "" {
		return NewValidationError("device_id", "设备ID不能为空")
	}
	if s.Direction == "" {
		return NewValidationError("direction", "同步方向不能为空")
	}
	if s.Direction != SyncLogDirectionPush && s.Direction != SyncLogDirectionPull {
		return NewValidationError("direction", "同步方向不合法")
	}
	if s.Status == "" {
		s.Status = SyncLogStatusSuccess
	}
	if s.Status != SyncLogStatusSuccess && s.Status != SyncLogStatusFailed {
		return NewValidationError("status", "同步状态不合法")
	}
	if s.EntryCount < 0 {
		return NewValidationError("entry_count", "条目数不能为负数")
	}
	s.Message = strings.TrimSpace(s.Message)
	return nil
}

type SyncLogFilter struct {
	DeviceID  string
	Direction string
	Status    string
}

func (f SyncLogFilter) Match(s *SyncLog) bool {
	if f.DeviceID != "" && s.DeviceID != f.DeviceID {
		return false
	}
	if f.Direction != "" && s.Direction != f.Direction {
		return false
	}
	if f.Status != "" && s.Status != f.Status {
		return false
	}
	return true
}
