package remote

import "fmt"

type Failure struct {
	Temporary bool
	Message   string
}

func (e *Failure) Error() string { return e.Message }

func Normalize(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("relay rejected: %v", err)
}
