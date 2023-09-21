package errorext

import "fmt"

type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}

func NewNotFoundError(format string, a ...any) *NotFoundError {
	return &NotFoundError{
		Message: fmt.Sprintf(format, a...),
	}
}
