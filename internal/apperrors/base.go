package apperrors

type AppError struct {
	HTTPStatusCode int    `json:"-"`
	Message        string `json:"message"`
	WrappedError   error  `json:"-"`
}

func (a *AppError) Error() string {
	msg := a.Message

	if a.WrappedError != nil {
		msg += " - Wrapped Error: " + a.WrappedError.Error()
	}

	return msg
}
