package apperr

import "fmt"

func Domain(code, message string) *AppError {
	return &AppError{
		Type:    DomainError,
		Code:    code,
		Message: message,
	}
}

func InvalidStatusTransition(from, to string) *AppError {
	return Domain(
		"invalid_status_transition",
		fmt.Sprintf("cannot transition from %s to %s", from, to),
	)
}

func InvalidFieldFormat(field, reason string) *AppError {
	return Domain(
		"invalid_field_format",
		fmt.Sprintf("invalid field format for %s: %s", field, reason),
	)
}
