package params

import "fmt"


type ThirdFilterValidator struct {
	IsRequired bool
}

func (tfv ThirdFilterValidator) IsValid(value string) error {
	if tfv.IsRequired && value == "" {
		return fmt.Errorf("is required")
	}

	return nil
}

func (tfv ThirdFilterValidator) KeyParam() string {
	return "third"
}