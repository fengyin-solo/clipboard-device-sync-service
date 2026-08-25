package syncflow

import "clipboard/internal/fanout"

type FanoutService struct {
	producer fanout.Producer
}

func NewFanoutService(producer fanout.Producer) *FanoutService {
	return &FanoutService{producer: producer}
}

func (s *FanoutService) Collect(values []string) ([]string, error) {
	resultsCh, errorsCh := s.producer.Produce(values)
	var results []string
	for value := range resultsCh {
		results = append(results, value)
	}
	if err := <-errorsCh; err != nil {
		return results, err
	}
	return results, nil
}
