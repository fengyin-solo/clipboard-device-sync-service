package syncflow

import "clipboard/internal/version"

type VersionService struct {
	store *version.Store
}

func NewVersionService(store *version.Store) *VersionService {
	return &VersionService{store: store}
}

func (s *VersionService) Complete(key string, versionNumber int) bool {
	return s.store.Put(key, version.State{Version: versionNumber, Status: "complete"})
}

func (s *VersionService) DelayedProgress(key string, versionNumber int) bool {
	return s.store.Put(key, version.State{Version: versionNumber, Status: "syncing"})
}

func (s *VersionService) State(key string) version.State { return s.store.Get(key) }
