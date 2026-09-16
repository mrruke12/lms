package enrollment

import (
	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/apperr"
)

type Enrollment struct {
	id       uuid.UUID
	userID   uuid.UUID
	objectID uuid.UUID
	typ      Type
	status   Status
}

func NewEnrollment(userID, objectID uuid.UUID, typ Type) (*Enrollment, error) {
	if !enrollmentTypeSet.Has(typ) {
		return nil, apperr.ConstraintViolation("ObjectID", "")
	}
}

/*
Setters
*/

func (e *Enrollment) SetStatus(status Status) error {
	if !enrollmentStatusTransitions.CanTransition(e.status, e.status) {
		return apperr.InvalidStatusTransition(string(e.status), string(status))
	}

	e.status = status

	return nil
}

/*
Getters
*/

func (e *Enrollment) ID() uuid.UUID {
	return e.id
}

func (e *Enrollment) Type() Type {
	return e.typ
}

func (e *Enrollment) Status() Status {
	return e.status
}
