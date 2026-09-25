package domainerr

import (
	"errors"

	erraggregation "github.com/mrruke12/lms/pkg/erragregation"
)

var (
	ErrInvalidFieldValue       error = errors.New("invalid field format")
	ErrField                   error = errors.New("constraint violation")
	ErrInvalidStatus           error = errors.New("invalid status")
	ErrInvalidStatusTransition error = errors.New("invalid status transition")
)

func NewInvalidFieldValue(field, reason string) error {
	return erraggregation.Join(ErrInvalidFieldValue, field, reason)
}

func NewInvalidStatus[T ~string](status T) error {
	return erraggregation.Join(ErrInvalidStatus, status)
}

func NewInvalidStatusTransition[T ~string](from, to T) error {
	return erraggregation.Join(ErrInvalidStatusTransition, from+" -> "+to)
}
