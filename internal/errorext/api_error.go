package errorext

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

type ApiError struct {
	Err        error  `json:"-"`
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

func (e ApiError) Error() string {
	return e.Err.Error()
}

func (e ApiError) StackTrace() string {
	if stackErr, ok := e.Err.(stackTracer); ok {
		stackTrace := stackErr.StackTrace()

		return fmt.Sprintf("%+v", stackTrace)
	}

	return "the stack trace is not exists for the error"
}

func ToApiError(err error) *ApiError {
	if err == nil {
		return nil
	}

	message := err.Error()
	statusCode := 500

	switch e := errors.Cause(err).(type) {
	case *ValidationError:
		statusCode = http.StatusBadRequest
		message = e.Error()
	case *NotFoundError:
		statusCode = http.StatusNotFound
		message = e.Error()
	}

	return &ApiError{
		Err:        err,
		Message:    message,
		StatusCode: statusCode,
	}
}
