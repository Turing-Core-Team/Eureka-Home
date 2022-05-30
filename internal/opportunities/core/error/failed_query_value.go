package error

import "fmt"

type FailedQueryValue struct {
	Message string
	Err     error
}

func NewFailedQueryValue(field, value string) FailedQueryValue {
	return FailedQueryValue{
		Message: fmt.Sprintf(
			`field:%s_value%s`,
			field,
			value,
		),
	}
}

func (uc FailedQueryValue) Error() string {
	return uc.Message
}
