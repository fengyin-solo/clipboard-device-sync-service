package syncflow

import (
	"testing"

	"clipboard/internal/catalog"
)

func TestCatalogCaptureRemainsImmutable(t *testing.T) {
	store := catalog.NewStore()
	store.Put("first", "alpha")
	service := NewCatalogService(store)
	captured := service.Capture()
	captured["first"] = "caller-change"
	store.Put("second", "bravo")
	last := service.LastCapture()
	if last["first"] != "alpha" {
		t.Fatalf("saved capture was changed by caller to %q", last["first"])
	}
	if _, ok := last["second"]; ok {
		t.Fatal("saved capture gained an item written after capture")
	}
}
