package apperr

import (
	"errors"

	"github.com/mrruke12/lms/internal/domains/domainerr"
)

func MapDBErr(err error) error {
	switch {
	case errors.Is(err, domainerr.ErrNotFound):
		return Error("storage_not_found", err.Error())
	case errors.Is(err, domainerr.ErrConflict):
		return Error("storage_conflict", err.Error())
	case errors.Is(err, domainerr.ErrForbidden):
		return Error("storage_forbidden", err.Error())
	case errors.Is(err, domainerr.ErrUnaviable):
		return Error("storage_unaviable", err.Error())
	case errors.Is(err, domainerr.ErrTimeout):
		return Error("storage_timeout", err.Error())
	default:
		return Wrap("storage_error", "unexpected storage error", err)
	}
}
