package syncflow

import (
	"clipboard/internal/frame"
	"clipboard/internal/snapshot"
)

type FrameService struct {
	decoder *frame.Decoder
	cache   *snapshot.Cache
}

func NewFrameService(decoder *frame.Decoder, cache *snapshot.Cache) *FrameService {
	return &FrameService{decoder: decoder, cache: cache}
}

func (s *FrameService) Store(key, text string) frame.Frame {
	decoded := s.decoder.Decode(text)
	s.cache.Put(key, decoded.Payload)
	return decoded
}

func (s *FrameService) Load(key string) []byte {
	return s.cache.Get(key)
}
