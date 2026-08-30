package lesson

import (
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/apperr"
)

type Lesson struct {
	id       uuid.UUID
	authorID uuid.UUID
	version  uuid.UUID
	name     string
	status   Status

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	UpdatedBy uuid.UUID
	DeletedBy *uuid.UUID
}

func NewLesson(authorID uuid.UUID, name string) *Lesson {
	return &Lesson{
		name:   name,
		status: StatusDraft,
	}
}

/*
Domain rules
*/

var isTrimmedRegexp = regexp.MustCompile(`^\S.*\S$`)

const nameMinLen = 6
const nameMaxLen = 128

func (l *Lesson) CanEdit() bool {
	if l.status == StatusPublished || l.status == StatusArchived {
		return false
	}

	return true
}

/*
Setters
*/

func (l *Lesson) SetStatus(status Status) error {
	if !IsValidStatus(status) {
		return apperr.InvalidStatus(string(status))
	}

	if !statusTransitions.CanTransition(l.status, status) {
		return apperr.InvalidStatusTransition(string(l.status), string(status))
	}

	return nil
}

func (l *Lesson) SetName(name string) error {
	if !isTrimmedRegexp.MatchString(name) {
		return apperr.InvalidFieldFormat("Name", "must start and end with non-whitespace character")
	}

	length := len(name)

	if length < nameMinLen {
		return apperr.InvalidFieldFormat("Name", "too short")
	}

	if length > nameMaxLen {
		return apperr.InvalidFieldFormat("Name", "too long")
	}

	return nil
}

/*
Getters
*/

func (l *Lesson) Version() uuid.UUID {
	return l.version
}

func (l *Lesson) ID() uuid.UUID {
	return l.id
}

func (l *Lesson) Name() string {
	return l.name
}

func (l *Lesson) Status() Status {
	return l.status
}
