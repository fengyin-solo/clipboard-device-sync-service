package resource

import "sync"

type Tracker struct {
	mu     sync.Mutex
	open   int
	limit  int
	closed int
}

func NewTracker(limit int) *Tracker { return &Tracker{limit: limit} }

func (t *Tracker) Open() (*Handle, error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.open >= t.limit {
		return nil, ErrLimit
	}
	t.open++
	return &Handle{tracker: t}, nil
}

func (t *Tracker) Counts() (int, int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.open, t.closed
}

type Handle struct {
	tracker *Tracker
	once    sync.Once
}

func (h *Handle) Close() {
	h.once.Do(func() {
		h.tracker.mu.Lock()
		defer h.tracker.mu.Unlock()
		h.tracker.open--
		h.tracker.closed++
	})
}
