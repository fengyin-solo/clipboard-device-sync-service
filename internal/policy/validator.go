package policy

import "fmt"

type Validator interface {
	Validate(string) error
}

type strictValidator struct {
	enabled bool
}

func (v *strictValidator) Validate(value string) error {
	if !v.enabled {
		return fmt.Errorf("validator disabled")
	}
	if value == "panic" {
		panic("validator failed")
	}
	if value == "" {
		return fmt.Errorf("empty clipboard content")
	}
	return nil
}

func NewValidator(enabled bool) Validator {
	if !enabled {
		return nil
	}
	return &strictValidator{enabled: true}
}
