package params

import "fmt"

type SecondFilterValidator struct {
	IsRequired bool
}

func (sfv SecondFilterValidator) IsValid(value string) error {
	if sfv.IsRequired && value == "" {
		return fmt.Errorf("is required")
	}

	return nil
}

func (sfv SecondFilterValidator) KeyParam() string {
	return "second"
}