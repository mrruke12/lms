package element

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Element struct {
	id         uuid.UUID
	lessonID   uuid.UUID
	parentID   *uuid.UUID
	typ        string
	assessment AssessmentType
	config     json.RawMessage

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewElement(
	lessonID uuid.UUID,
	parentID *uuid.UUID,
	typ string,
	assessment AssessmentType,
	config json.RawMessage,
) *Element {
	return &Element{
		lessonID:   lessonID,
		parentID:   parentID,
		typ:        typ,
		assessment: assessment,
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
	if err := validateJSON(elementSchema, config); err != nil {
		return err
	}

	if e.assessment != Assessment.None {
		if err := validateJSON(assessmentSchema, config); err != nil {
			return err
		}
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
		typ:        e.typ,
		assessment: e.assessment,
		config:     e.ConfigRaw(),
	}
}
