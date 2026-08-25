package syncflow

import (
	"fmt"

	"clipboard/internal/policy"
)

type RecordService struct {
	validator policy.Validator
	records   map[string]string
}

func NewRecordService(validator policy.Validator) *RecordService {
	return &RecordService{validator: validator, records: make(map[string]string)}
}

func (s *RecordService) Record(id, value string) (err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			err = fmt.Errorf("validation panic: %v", recovered)
		}
	}()
	if s.validator != nil {
		if err := s.validator.Validate(value); err != nil {
			return err
		}
	}
	s.records[id] = value
	return nil
}

func (s *RecordService) Get(id string) (string, bool) {
	value, ok := s.records[id]
	return value, ok
}
