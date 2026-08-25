package fanout

import "fmt"

type Producer struct{}

func (Producer) Produce(values []string) (<-chan string, <-chan error) {
	results := make(chan string)
	errorsCh := make(chan error, 1)
	go func() {
		// Always close results so the collector's range terminates, even on the
		// reject path where we return early — otherwise the collector blocks
		// forever waiting for more values and the rejection never surfaces.
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
