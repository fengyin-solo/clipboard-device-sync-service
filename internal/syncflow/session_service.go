package syncflow

import "clipboard/internal/session"

type SessionService struct {
	pool *session.Pool
}

func NewSessionService(pool *session.Pool) *SessionService {
	return &SessionService{pool: pool}
}

func (s *SessionService) Snapshot(deviceID string, payload []byte) func() session.Value {
	value := s.pool.Acquire(deviceID, payload)
	snapshot := session.Value{DeviceID: value.DeviceID, Payload: append([]byte(nil), value.Payload...)}
	s.pool.Release(value)
	return func() session.Value { return snapshot }
}
