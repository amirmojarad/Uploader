package errorext

import "fmt"

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

func NewValidationError(format string, a ...any) *ValidationError {
	return &ValidationError{
		Message: fmt.Sprintf(format, a...),
	}
}
