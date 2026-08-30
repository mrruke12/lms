package element

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Revision struct {
	elementID  uuid.UUID // Backward compatibility with element
	id         uuid.UUID
	lessonID   uuid.UUID
	parentID   *uuid.UUID
	typ        string
	assessment AssessmentType
	config     json.RawMessage

	CreatedAt time.Time
}

func NewRevision(
	lessonID uuid.UUID,
	parentID *uuid.UUID,
	typ string,
	assessment AssessmentType,
	config json.RawMessage,
) *Revision {
	return &Revision{
		lessonID:   lessonID,
		parentID:   parentID,
		typ:        typ,
		assessment: assessment,
		config:     config,
	}
}

/*
Getters
*/
func (e *Revision) ElementID() uuid.UUID {
	return e.elementID
}

func (e *Revision) ID() uuid.UUID {
	return e.id
}

func (e *Revision) LessonID() uuid.UUID {
	return e.lessonID
}

func (e *Revision) ParentID() *uuid.UUID {
	return e.parentID
}

func (e *Revision) Type() string {
	return e.typ
}

func (e *Revision) Assessment() AssessmentType {
	return e.assessment
}

func (e *Revision) ConfigRaw() json.RawMessage {
	config := make(json.RawMessage, len(e.config))
	copy(config, e.config)
	return config
}

func (e *Revision) ConfigUnmarshal(dst any) error {
	return json.Unmarshal(e.config, dst)
}
