package apperr

import (
	"fmt"
)

func Domain(code, message string) *AppError {
	return &AppError{
		Type:    DomainError,
		Code:    code,
		Message: message,
	}
}

func InvalidStatus(status string) *AppError {
	return Domain(
		"invalid_status",
		fmt.Sprintf("invalid status: %s", status),
	)
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

func InvalidJSONSchema(reason string) *AppError {
	return Domain(
		"invalid_json_schema",
		reason,
	)
}

func ConstraintViolation(constraint, violation string) *AppError {
	return Domain(
		"constraint_violation",
		fmt.Sprintf("constraint violated for %s: %s", violation),
	)
}
