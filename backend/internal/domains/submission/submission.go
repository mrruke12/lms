package submission

import (
	"time"

	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/apperr"
)

type Submission struct {
	id        uuid.UUID
	attemptID uuid.UUID
	elementID uuid.UUID
	cap       int
	grade     int
	status    Status

	CreatedAt time.Time
}

func NewSubmission(attemptID, elementID uuid.UUID, cap int) *Submission {
	return &Submission{
		attemptID: attemptID,
		elementID: elementID,
		cap:       cap,
		status:    StatusPending,
	}
}

/*
Setters
*/

func (s *Submission) SetStatus(status Status) error {
	if !IsValidStatus(status) {
		return apperr.InvalidStatus(string(status))
	}

	if !statusTransitions.CanTransition(s.status, status) {
		return apperr.InvalidStatusTransition(string(s.status), string(status))
	}

	s.status = status

	return nil
}

func (s *Submission) SetGrade(grade int) error {
	if grade < 0 {
		return apperr.ConstraintViolation("Grade", "cannot be negative")
	}

	if grade > s.cap {
		return apperr.ConstraintViolation("Grade", "cannot be greater than cap")
	}

	s.grade = grade

	return nil
}

/*
Getters
*/

func (s *Submission) ID() uuid.UUID {
	return s.id
}

func (s *Submission) AttemptID() uuid.UUID {
	return s.attemptID
}

func (s *Submission) ElementID() uuid.UUID {
	return s.elementID
}

func (s *Submission) Cap() int {
	return s.cap
}

func (s *Submission) Grade() int {
	return s.grade
}

func (s *Submission) Status() Status {
	return s.status
}
