package syncflow

import (
	"sync"

	"clipboard/internal/catalog"
)

type CatalogService struct {
	store *catalog.Store
	mu    sync.RWMutex
	last  map[string]string
}

func NewCatalogService(store *catalog.Store) *CatalogService {
	return &CatalogService{store: store}
}

func (s *CatalogService) Capture() map[string]string {
	snapshot := cloneCatalog(s.store.Snapshot())
	s.mu.Lock()
	s.last = snapshot
	s.mu.Unlock()
	// 返回独立副本，确保调用方拿到后与内部缓存互不影响。
	return cloneCatalog(snapshot)
}

func (s *CatalogService) LastCapture() map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.last == nil {
		return nil
	}
	// 返回独立副本，保证历史快照不可变。
	return cloneCatalog(s.last)
}

func cloneCatalog(source map[string]string) map[string]string {
	cloned := make(map[string]string, len(source))
	for key, value := range source {
		cloned[key] = value
	}
	return cloned
}
