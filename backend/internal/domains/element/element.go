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
	typeID     uuid.UUID
	assessment AssessmentType
	config     json.RawMessage

	CreatedAt time.Time
	UpdatedAt time.Time
}

/*
Setters
*/
func (e *Element) SetParentID(id *uuid.UUID) error {
	if id != nil && e.id == *id {
		return apperr.ConstraintViolation("ParentID", "cannot be parent of itself")
	}

	e.parentID = id

	return nil
}

func (e *Element) SetConfig(config json.RawMessage) error {
	if !json.Valid(config) {
		return apperr.InvalidJSONSchema("invalid json")
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

func (e *Element) TypeID() uuid.UUID {
	return e.typeID
}

func (e *Element) Assessment() AssessmentType {
	return e.assessment
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
		typeID:     e.typeID,
		assessment: e.assessment,
		config:     e.ConfigRaw(),
	}
}
