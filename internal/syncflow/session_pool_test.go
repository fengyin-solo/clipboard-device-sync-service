package syncflow

import (
	"testing"

	"clipboard/internal/session"
)

func TestAsyncSnapshotIsolatedFromSessionReuse(t *testing.T) {
	pool := session.NewPool()
	service := NewSessionService(pool)
	readFirst := service.Snapshot("device-a", []byte("alpha"))
	second := pool.Acquire("device-b", []byte("bravo"))
	pool.Release(second)
	first := readFirst()
	if first.DeviceID != "device-a" || string(first.Payload) != "alpha" {
		t.Fatalf("first async snapshot became device=%q payload=%q", first.DeviceID, first.Payload)
	}
	empty := pool.Acquire("", nil)
	if empty.DeviceID != "" || len(empty.Payload) != 0 {
		t.Fatalf("reused empty session kept device=%q payload=%q", empty.DeviceID, empty.Payload)
	}
}
