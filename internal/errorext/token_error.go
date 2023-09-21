package errorext

import "fmt"

type TokenError struct {
	Message string
}

func (e TokenError) Error() string {
	return e.Message
}

func NewTokenError(format string, a ...any) *TokenError {
	return &TokenError{
		Message: fmt.Sprintf(format, a...),
	}
}
