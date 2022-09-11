package params

import "fmt"

type AreaFilterValidator struct {
	IsRequired bool
}

func (tfv AreaFilterValidator) IsValid(value string) error {
	if tfv.IsRequired && value == "" {
		return fmt.Errorf("is required")
	}

	return nil
}

func (tfv AreaFilterValidator) KeyParam() string {
	return "area"
}
