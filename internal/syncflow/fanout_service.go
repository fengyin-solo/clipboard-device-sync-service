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
	for resultsCh != nil || errorsCh != nil {
		select {
		case value, ok := <-resultsCh:
			if !ok {
				resultsCh = nil
				continue
			}
			results = append(results, value)
		case err, ok := <-errorsCh:
			if !ok {
				errorsCh = nil
				continue
			}
			if err != nil {
				return results, err
			}
		}
	}
	return results, nil
}
