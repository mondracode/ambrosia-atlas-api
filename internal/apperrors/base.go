package apperrors

type AppError struct {
	HTTPStatusCode int    `json:"-"`
	Message        string `json:"message"`
	WrappedError   error  `json:"-"`
}

func NewAppError(httpStatusCode int, message string, wrappedError error) *AppError {
	return &AppError{
		HTTPStatusCode: httpStatusCode,
		Message:        message,
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
