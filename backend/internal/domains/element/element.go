package element

import (
	"encoding/json"
	"errors"
	"slices"

	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/domains/assessment"
	"github.com/mrruke12/lms/internal/domains/domainerr"
)

type Element struct {
	id         uuid.UUID
	lessonID   uuid.UUID
	parentID   *uuid.UUID
	typeID     int
	assessment assessment.Type
	config     json.RawMessage
}

/*
Setters
*/
func (e *Element) SetParentID(id *uuid.UUID) error {
	if id != nil && e.id == *id {
		return domainerr.NewInvalidFieldValue("ParentID", "cannot be parent of itself")
	}

	e.parentID = id

	return nil
}

func (e *Element) SetConfig(config json.RawMessage) error {
	if !json.Valid(config) {
		return errors.New("invalid json")
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

func (e *Element) TypeID() int {
	return e.typeID
}

func (e *Element) Assessment() assessment.Type {
	return e.assessment
}

func (e *Element) ConfigRaw() json.RawMessage {
	return slices.Clone(e.config)
}

func (e *Element) ConfigUnmarshal(dst any) error {
	return json.Unmarshal(e.config, dst)
}
