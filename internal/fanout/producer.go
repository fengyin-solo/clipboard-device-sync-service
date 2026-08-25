package fanout

import "fmt"

type Producer struct{}

func (Producer) Produce(values []string) (<-chan string, <-chan error) {
	results := make(chan string)
	errorsCh := make(chan error, 1)
	go func() {
		defer close(results)
		defer close(errorsCh)
		for _, value := range values {
			if value == "reject" {
				errorsCh <- fmt.Errorf("rejected sync value")
				return
			}
			results <- value
		}
	}()
	return results, errorsCh
}
