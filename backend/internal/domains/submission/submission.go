package submission

import (
	"time"

	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/domains/domainerr"
)

type Submission struct {
	id        uuid.UUID
	attemptID uuid.UUID
	elementID uuid.UUID
	cap       int
	score     int
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
		return domainerr.NewInvalidStatus(string(status))
	}

	if !statusTransitions.CanTransition(s.status, status) {
		return domainerr.NewInvalidStatusTransition(string(s.status), string(status))
	}

	s.status = status

	return nil
}

func (s *Submission) SetScore(score int) error {
	if score < 0 {
		return domainerr.NewInvalidFieldValue("Grade", "cannot be negative")
	}

	if score > s.cap {
		return domainerr.NewInvalidFieldValue("Grade", "cannot be greater than cap")
	}

	s.score = score

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

func (s *Submission) Score() int {
	return s.score
}

func (s *Submission) Status() Status {
	return s.status
}
