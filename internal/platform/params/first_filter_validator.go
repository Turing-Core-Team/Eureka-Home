package params

import "fmt"

const (
	person  string = "personas"
	project string = "proyectos"
)

type WhoFilterValidator struct {
	IsRequired bool
}

func (ffv WhoFilterValidator) IsValid(value string) error {
	if ffv.IsRequired && value == "" {
		return fmt.Errorf("is required")
	}

	if value != person && value != project {
		return fmt.Errorf("bad request values who_filter")
	}

	return nil
}

func (ffv WhoFilterValidator) KeyParam() string {
	return "who"
}
