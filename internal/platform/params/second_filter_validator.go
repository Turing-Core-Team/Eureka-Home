package params

import "fmt"

type TypeFilterValidator struct {
	IsRequired bool
}

func (sfv TypeFilterValidator) IsValid(value string) error {
	if sfv.IsRequired && value == "" {
		return fmt.Errorf("is required")
	}

	return nil
}

func (sfv TypeFilterValidator) KeyParam() string {
	return "type"
}
