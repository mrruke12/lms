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
	status    Status

	CreatedAt time.Time
}

func NewSubmission(attemptID, elementID uuid.UUID) *Submission {
	return &Submission{
		attemptID: attemptID,
		elementID: elementID,
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

func (s *Submission) Status() Status {
	return s.status
}
