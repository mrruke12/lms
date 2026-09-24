package domainerr

import (
	"errors"

	erraggregation "github.com/mrruke12/lms/pkg/erragregation"
)

var (
	ErrInvalidFieldFormat      error = errors.New("invalid field format")
	ErrConstraintViolation     error = errors.New("constraint violation")
	ErrInvalidStatus           error = errors.New("invalid status")
	ErrInvalidStatusTransition error = errors.New("invalid status transition")
)

func NewInvalidFieldFormat(field, reason string) error {
	return erraggregation.Join(ErrInvalidFieldFormat, field, reason)
}

func NewConstraintViolation(constraint, reason string) error {
	return erraggregation.Join(ErrConstraintViolation, constraint, reason)
}

func NewInvalidStatus[T ~string](status T) error {
	return erraggregation.Join(ErrInvalidStatus, status)
}

func NewInvalidStatusTransition[T ~string](from, to T) error {
	return erraggregation.Join(ErrInvalidStatusTransition, from+" -> "+to)
}
