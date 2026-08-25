package store

import (
	"sync"

	"clipboard/internal/model"
)

type MemoryStore struct {
	mu        sync.RWMutex
	entries   map[string]*model.Entry
	groups    map[string]*model.Group
	tags      map[string]*model.Tag
	entryTags map[string]*model.EntryTag
	devices   map[string]*model.Device
	syncLogs  map[string]*model.SyncLog
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		entries:   make(map[string]*model.Entry),
		groups:    make(map[string]*model.Group),
		tags:      make(map[string]*model.Tag),
		entryTags: make(map[string]*model.EntryTag),
		devices:   make(map[string]*model.Device),
		syncLogs:  make(map[string]*model.SyncLog),
	}
}

var _ Store = (*MemoryStore)(nil)
