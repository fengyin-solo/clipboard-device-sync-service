// Package store 定义数据访问接口与内存实现。
package store

import (
	"errors"

	"clipboard/internal/model"
)

var (
	ErrNotFound = errors.New("记录不存在")
	ErrConflict = errors.New("记录已存在或状态冲突")
)

// Store 聚合全部实体的数据访问方法，便于测试时替换实现。
type Store interface {
	// Entry
	CreateEntry(e *model.Entry) error
	GetEntry(id string) (*model.Entry, error)
	ListEntries() []*model.Entry
	UpdateEntry(e *model.Entry) error
	DeleteEntry(id string) error

	// Group
	CreateGroup(g *model.Group) error
	GetGroup(id string) (*model.Group, error)
	ListGroups() []*model.Group
	UpdateGroup(g *model.Group) error
	DeleteGroup(id string) error

	// Tag
	CreateTag(t *model.Tag) error
	GetTag(id string) (*model.Tag, error)
	ListTags() []*model.Tag
	UpdateTag(t *model.Tag) error
	DeleteTag(id string) error

	// EntryTag
	CreateEntryTag(et *model.EntryTag) error
	GetEntryTag(id string) (*model.EntryTag, error)
	ListEntryTags() []*model.EntryTag
	DeleteEntryTag(id string) error
	ListEntryTagsByEntry(entryID string) []*model.EntryTag
	ListEntryTagsByTag(tagID string) []*model.EntryTag

	// Device
	CreateDevice(d *model.Device) error
	GetDevice(id string) (*model.Device, error)
	ListDevices() []*model.Device
	UpdateDevice(d *model.Device) error
	DeleteDevice(id string) error

	// SyncLog
	CreateSyncLog(s *model.SyncLog) error
	GetSyncLog(id string) (*model.SyncLog, error)
	ListSyncLogs() []*model.SyncLog
	UpdateSyncLog(s *model.SyncLog) error
	DeleteSyncLog(id string) error
}
