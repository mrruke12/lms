package apperr

import "fmt"

func Infrastructure(code, message string) *AppError {
	return &AppError{
		Type:    InfrastructureError,
		Code:    code,
		Message: message,
	}
}

func EnvironmentVariableError(reason string) *AppError {
	return Infrastructure(
		"environment_variable_error",
		fmt.Sprintf("environment variable error: %s", reason),
	)
}
