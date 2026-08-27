package lesson

import (
	"regexp"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/apperr"
)

type Lesson struct {
	id     uuid.UUID
	name   string
	status LessonStatus

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	CreatedBy uuid.UUID
	UpdatedBy uuid.UUID
	DeletedBy *uuid.UUID
}

func NewLesson(name string) *Lesson {
	return &Lesson{
		name:   name,
		status: Status.Draft,
	}
}

/*
Domain rules
*/

var statusTransitions = map[LessonStatus][]LessonStatus{
	Status.Draft: {
		Status.Published,
	},
	Status.Published: {
		Status.Editing,
	},
	Status.Editing: {
		Status.Published,
		Status.Archived,
	},
	Status.Archived: {
		Status.Draft,
	},
}

var isTrimmedRegexp = regexp.MustCompile(`^\S.*\S$`)

const nameMinLen = 6
const nameMaxLen = 128

func (l *Lesson) CanEdit() bool {
	if l.status == Status.Published || l.status == Status.Archived {
		return false
	}

	return true
}

/*
Setters
*/

func (l *Lesson) SetStatus(status LessonStatus) error {
	transitions := statusTransitions[l.status]

	if slices.Contains(transitions, status) == false {
		return apperr.InvalidStatusTransition(string(l.status), string(status))
	}

	return nil
}

func (l *Lesson) SetName(name string) error {
	if isTrimmedRegexp.MatchString(name) == false {
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

func (l *Lesson) ID() uuid.UUID {
	return l.id
}

func (l *Lesson) Name() string {
	return l.name
}

func (l *Lesson) Status() LessonStatus {
	return l.status
}
