package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mrruke12/lms/internal/domains/domainerr"
	"github.com/mrruke12/lms/pkg/erraggregation"
)

func wrapDriverErr(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, context.DeadlineExceeded) {
		return erraggregation.Join(domainerr.ErrTimeout, err)
	}
	if errors.Is(err, context.Canceled) {
		return erraggregation.Join(domainerr.ErrUnavailable, err) // или ErrTimeout — на твой вкус
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505": // unique_violation
			return erraggregation.Join(domainerr.ErrConflict, err)
		case "23503": // foreign_key_violation
			return erraggregation.Join(domainerr.ErrConflict, err)
		case "23502": // not_null_violation
			return erraggregation.Join(domainerr.ErrConflict, err)
		case "40001": // serialization_failure
			return erraggregation.Join(domainerr.ErrConflict, err)
		case "40P01": // deadlock_detected
			return erraggregation.Join(domainerr.ErrConflict, err)
		}
	}

	var connErr *pgconn.ConnectError
	if errors.As(err, &connErr) {
		return erraggregation.Join(domainerr.ErrUnavailable, err)
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return erraggregation.Join(domainerr.ErrNotFound, err)
	}

	return err
}
