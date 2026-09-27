package revision

import (
	"encoding/json"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/domains/element"
	"github.com/mrruke12/lms/pkg/storage"
)

type Revision struct {
	elementID  uuid.UUID // Backward compatibility with element
	id         uuid.UUID
	lessonID   uuid.UUID
	parentID   *uuid.UUID
	typ        string
	assessment string
	config     json.RawMessage

	CreatedAt time.Time
}

/*
Constructors
*/

func FromElement(el *element.Element) *Revision {
	return &Revision{
		elementID:  el.ID(),
		lessonID:   el.LessonID(),
		parentID:   el.ParentID(),
		typ:        string(el.Type()),
		assessment: string(el.Assessment()),
		config:     el.ConfigRaw(),
	}
}

func FromRow(scanner storage.Scanner) (*Revision, error) {
	rev := &Revision{}

	err := scanner.Scan(
		&rev.elementID,
		&rev.id,
		&rev.lessonID,
		&rev.parentID,
		&rev.typ,
		&rev.assessment,
		&rev.config,
	)

	if err != nil {
		return nil, err
	}

	return rev, nil
}

/*
Getters
*/

func (r *Revision) ElementID() uuid.UUID {
	return r.elementID
}

func (r *Revision) ID() uuid.UUID {
	return r.id
}

func (r *Revision) LessonID() uuid.UUID {
	return r.lessonID
}

func (r *Revision) ParentID() *uuid.UUID {
	return r.parentID
}

func (r *Revision) Type() string {
	return r.typ
}

func (r *Revision) Assessment() string {
	return r.assessment
}

func (r *Revision) ConfigRaw() json.RawMessage {
	return slices.Clone(r.config)
}

func (r *Revision) ConfigUnmarshal(dst any) error {
	return json.Unmarshal(r.config, dst)
}

/*
Helpers
*/

// Reports whether the element differs from revision
func (r *Revision) Differs(el *element.Element) bool {
	return el.ParentID() != r.ParentID() ||
		!slices.Equal(r.config, el.ConfigRaw())
}
