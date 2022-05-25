package params

import "fmt"

const (
	person string = "person"
	project string = "proyecto"
)

type FirstFilterValidator struct {
	IsRequired bool
}

func (ffv FirstFilterValidator) IsValid(value string) error {
	if ffv.IsRequired && value == "" {
		return fmt.Errorf("is required")
	}

	if value != person || value != project {
		return fmt.Errorf("bad request values first filter")
	}

	return nil
}

func (ffv FirstFilterValidator) KeyParam() string {
	return "First Filter"
}

