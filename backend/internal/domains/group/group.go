package group

import (
	"time"

	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/domains/domainerr"
)

type Group struct {
	id             uuid.UUID
	name           string
	degree         Degree
	entryYear      int
	graduationYear int
}

func NewGroup(name string, degree Degree, entryYear, graduationYear int) (*Group, error) {
	if valid, reason := isValidName(name); !valid {
		return nil, domainerr.NewConstraintViolation("Name", reason)
	}

	if entryYear > time.Now().Year() {
		return nil, domainerr.NewConstraintViolation("EntryYear", "cannot be a future year")
	}

	if entryYear < graduationYear {
		return nil, domainerr.NewConstraintViolation("EntryYear", "cannot be before graduation year")
	}

	return &Group{
		name:           name,
		degree:         degree,
		entryYear:      entryYear,
		graduationYear: graduationYear,
	}, nil
}

/*
Domain rules
*/

func isValidName(name string) (bool, string) {
	ln := len(name)

	if ln < 3 {
		return false, "too short"
	}

	if ln > 256 {
		return false, "too long"
	}

	return true, ""
}

/*
Setters
*/

func (g *Group) SetName(name string) error {
	if valid, reason := isValidName(name); !valid {
		return domainerr.NewConstraintViolation("Name", reason)
	}

	g.name = name

	return nil
}

/*
Getters
*/

func (g *Group) ID() uuid.UUID {
	return g.id
}

func (g *Group) Name() string {
	return g.name
}

func (g *Group) Degree() Degree {
	return g.degree
}

func (g *Group) EntryYear() int {
	return g.entryYear
}

func (g *Group) GraduationYear() int {
	return g.graduationYear
}
