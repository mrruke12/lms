package assessment

import (
	"time"

	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/apperr"
)

type Assessment struct {
	id        uuid.UUID
	attemptID uuid.UUID
	status    Status
	grade     int16

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAssessment(attemptID uuid.UUID) *Assessment {
	return &Assessment{
		attemptID: attemptID,
		grade:     0,
	}
}

/*
Setters
*/

func (a *Assessment) SetStatus(status Status) error {
	if !statusSet.Has(status) {
		return apperr.InvalidStatus(string(status))
	}

	if !statusTransitions.CanTransition(a.status, status) {
		return apperr.InvalidStatusTransition(string(a.status), string(status))
	}

	a.status = status

	return nil
}

func (a *Assessment) SetGrade(grade int16) error {
	if grade < 0 || grade > 100 {
		return apperr.InvalidFieldFormat("Grade", "must be in range 0-100")
	}

	a.grade = grade

	return nil
}

/*
Getters
*/

func (a *Assessment) ID() uuid.UUID {
	return a.id
}

func (a *Assessment) AttemptID() uuid.UUID {
	return a.attemptID
}

func (a *Assessment) Status() Status {
	return a.status
}

func (a *Assessment) Grade() int16 {
	return a.grade
}
