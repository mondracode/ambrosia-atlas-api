package apperrors

import "net/http"

type AppError struct {
	HTTPStatusCode int
	Message        string
	ErrorType      string
	WrappedError   error
}

func NewAppError(httpStatusCode int, message string, wrappedError error) *AppError {
	errorType := http.StatusText(httpStatusCode)
	if errorType == "" {
		errorType = "Unknown"
	}

	return &AppError{
		HTTPStatusCode: httpStatusCode,
		Message:        message,
		ErrorType:      errorType,
		WrappedError:   wrappedError,
	}
}

func (a AppError) Error() string {
	msg := a.Message

	if a.WrappedError != nil {
		msg += " - Wrapped Error: " + a.WrappedError.Error()
	}

	return msg
}
