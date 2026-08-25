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
	snapshot := s.store.Snapshot()
	s.mu.Lock()
	s.last = snapshot
	s.mu.Unlock()
	return snapshot
}

func (s *CatalogService) LastCapture() map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.last
}

func cloneCatalog(source map[string]string) map[string]string {
	cloned := make(map[string]string, len(source))
	for key, value := range source {
		cloned[key] = value
	}
	return cloned
}
