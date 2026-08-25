package syncflow

import (
	"testing"

	"clipboard/internal/document"
)

func TestRecoveredBuildDoesNotCachePartialDocument(t *testing.T) {
	service := NewDocumentService(document.Builder{}, document.NewCache())
	if err := service.Build("doc-1", "panic"); err == nil {
		t.Fatal("panic build did not return an error")
	}
	if partial, ok := service.Get("doc-1"); ok {
		t.Fatalf("panic build left cached partial document: %+v", partial)
	}
	if err := service.Build("doc-1", "ready"); err != nil {
		t.Fatalf("next valid build failed: %v", err)
	}
	ready, ok := service.Get("doc-1")
	if !ok || !ready.Ready || ready.Payload != "ready" {
		t.Fatalf("valid rebuild returned document=%+v found=%v", ready, ok)
	}
}
