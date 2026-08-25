package catalog

import "sync"

type Store struct {
	mu    sync.RWMutex
	items map[string]string
}

func NewStore() *Store { return &Store{items: make(map[string]string)} }

func (s *Store) Put(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[key] = value
}

func (s *Store) Snapshot() map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	// 返回内部 map 的深拷贝，避免调用方对返回值的写操作、
	// 或后续 Put 影响到已生成的快照。快照一旦生成即应不可变。
	snapshot := make(map[string]string, len(s.items))
	for key, value := range s.items {
		snapshot[key] = value
	}
	return snapshot
}
