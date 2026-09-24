package lesson

import (
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/domains/domainerr"
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

func NewLesson(authorID uuid.UUID, name string) (*Lesson, error) {
	if err := isValidName(name); err != nil {
		return nil, err
	}

	return &Lesson{
		name:   name,
		status: StatusDraft,
	}, nil
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

func isValidName(name string) error {
	if !isTrimmedRegexp.MatchString(name) {
		return domainerr.NewInvalidFieldFormat("Name", "must start and end with non-whitespace character")
	}

	length := len(name)

	if length < nameMinLen {
		return domainerr.NewInvalidFieldFormat("Name", "too short")
	}

	if length > nameMaxLen {
		return domainerr.NewInvalidFieldFormat("Name", "too long")
	}

	return nil
}

/*
Setters
*/

func (l *Lesson) SetStatus(status Status) error {
	if !IsValidStatus(status) {
		return domainerr.NewInvalidStatus(string(status))
	}

	if !statusTransitions.CanTransition(l.status, status) {
		return domainerr.NewInvalidStatusTransition(string(l.status), string(status))
	}

	return nil
}

func (l *Lesson) SetName(name string) error {
	if err := isValidName(name); err != nil {
		return err
	}

	l.name = name

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
