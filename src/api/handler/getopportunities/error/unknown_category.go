package error

import "fmt"

type UnknownCategory struct {
	message string
}

func NewUnknownCategory(field, value string) UnknownCategory {
	return UnknownCategory{
		message: fmt.Sprintf(
			`field:%s_value%s`,
			field,
			value,
		),
	}
}

func (uc UnknownCategory) Error() string {
	return uc.message
}