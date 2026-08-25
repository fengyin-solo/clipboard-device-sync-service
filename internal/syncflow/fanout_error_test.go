package syncflow

import (
	"testing"
	"time"

	"clipboard/internal/fanout"
)

func TestFanoutStopsOnRejectedValue(t *testing.T) {
	service := NewFanoutService(fanout.Producer{})
	done := make(chan error, 1)
	go func() {
		_, err := service.Collect([]string{"first", "reject"})
		done <- err
	}()
	select {
	case err := <-done:
		if err == nil {
			t.Fatal("rejected fanout returned without its rejection error")
		}
	case <-time.After(250 * time.Millisecond):
		t.Fatal("rejected fanout did not return after producer stopped")
	}
}
