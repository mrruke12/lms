package apperr

func Application(code, message string) *AppError {
	return &AppError{
		Type:    ApplicationError,
		Code:    code,
		Message: message,
	}
}

func Forbidden() *AppError {
	return Application(
		"forbidden",
		"access denied",
	)
}
