package element

import (
	"encoding/json"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/domains/element"
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
