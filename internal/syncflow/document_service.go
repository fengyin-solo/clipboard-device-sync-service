package syncflow

import "clipboard/internal/document"

type DocumentService struct {
	builder document.Builder
	cache   *document.Cache
}

func NewDocumentService(builder document.Builder, cache *document.Cache) *DocumentService {
	return &DocumentService{builder: builder, cache: cache}
}

func (s *DocumentService) Build(id, payload string) error {
	doc, err := s.builder.Build(id, payload)
	s.cache.Put(doc)
	if err != nil {
		return err
	}
	return nil
}

func (s *DocumentService) Get(id string) (*document.Document, bool) {
	return s.cache.Get(id)
}
