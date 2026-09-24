package attempt

import (
	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/domains/domainerr"
)

type Attempt struct {
	id            uuid.UUID
	lessonID      uuid.UUID
	lessonVersion uuid.UUID
	studentID     uuid.UUID
	status        Status
}

func NewAttempt(lessonID, lessonVersion, studentID uuid.UUID) *Attempt {
	return &Attempt{
		lessonID:      lessonID,
		lessonVersion: lessonVersion,
		studentID:     studentID,
		status:        StatusActive,
	}
}

/*
Setters
*/

func (a *Attempt) SetStatus(status Status) error {
	if !IsValidStatus(status) {
		return domainerr.NewInvalidStatus(string(status))
	}

	if !statusTransitions.CanTransition(a.status, status) {
		return domainerr.NewInvalidStatusTransition(string(a.status), string(status))
	}

	a.status = status

	return nil
}

/*
Getters
*/

func (a *Attempt) ID() uuid.UUID {
	return a.id
}

func (a *Attempt) LessonID() uuid.UUID {
	return a.lessonID
}

func (a *Attempt) LessonVersion() uuid.UUID {
	return a.lessonVersion
}

func (a *Attempt) StudentID() uuid.UUID {
	return a.studentID
}

func (a *Attempt) Status() Status {
	return a.status
}
