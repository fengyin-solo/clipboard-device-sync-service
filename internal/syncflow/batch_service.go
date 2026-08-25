package syncflow

import "clipboard/internal/resource"

type BatchService struct {
	tracker *resource.Tracker
}

func NewBatchService(tracker *resource.Tracker) *BatchService {
	return &BatchService{tracker: tracker}
}

func (s *BatchService) Process(values []string) error {
	for range values {
		if err := s.processOne(); err != nil {
			return err
		}
	}
	return nil
}

func (s *BatchService) processOne() error {
	handle, err := s.tracker.Open()
	if err != nil {
		return err
	}
	defer handle.Close()
	return nil
}
