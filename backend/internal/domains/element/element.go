package element

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/apperr"
)

type Element struct {
	id         uuid.UUID
	lessonID   uuid.UUID
	parentID   *uuid.UUID
	typ        string
	assessable bool
	config     json.RawMessage

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	CreatedBy uuid.UUID
	UpdatedBy uuid.UUID
	DeletedBy *uuid.UUID
}

func NewElement(
	lessonID uuid.UUID,
	parentID *uuid.UUID,
	typ string,
	assessable bool,
	config json.RawMessage,
) *Element {
	return &Element{
		lessonID:   lessonID,
		parentID:   parentID,
		typ:        typ,
		assessable: assessable,
		config:     config,
	}
}

/*
Setters
*/

func (e *Element) SetParentID(id *uuid.UUID) {
	e.parentID = id
}

func (e *Element) SetConfig(config json.RawMessage) error {
	// TODO: define config by type and validate it
	if json.Valid(config) == false {
		return apperr.InvalidFieldFormat("Config", "invalid JSON")
	}

	e.config = config

	return nil
}

/*
Getters
*/

func (e *Element) ID() uuid.UUID {
	return e.id
}

func (e *Element) LessonID() uuid.UUID {
	return e.lessonID
}

func (e *Element) ParentID() *uuid.UUID {
	return e.parentID
}

func (e *Element) Type() string {
	return e.typ
}

func (e *Element) Assessable() bool {
	return e.assessable
}

func (e *Element) ConfigRaw() json.RawMessage {
	config := make(json.RawMessage, len(e.config))
	copy(config, e.config)
	return config
}

func (e *Element) ConfigUnmarshal(dst any) error {
	return json.Unmarshal(e.config, dst)
}

/*
Helpers
*/

func (e *Element) ToRevision() *Revision {
	return &Revision{
		elementID:  e.id,
		lessonID:   e.lessonID,
		parentID:   e.parentID,
		typ:        e.typ,
		assessable: e.assessable,
		config:     e.ConfigRaw(),
	}
}
