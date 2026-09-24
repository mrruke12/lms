package apperr

type Type string

type AppError struct {
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}

	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func Wrap(code, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func Error(code, message string) error {
	return &AppError{
		Code:    code,
		Message: message,
	}
}
