package element

import (
	"encoding/json"
	"errors"
	"slices"

	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/domains/assessment"
	"github.com/mrruke12/lms/internal/domains/domainerr"
	"github.com/mrruke12/lms/pkg/enum"
	erraggregation "github.com/mrruke12/lms/pkg/erragregation"
	"github.com/mrruke12/lms/pkg/storage"
)

type Element struct {
	id         uuid.UUID
	lessonID   uuid.UUID
	parentID   *uuid.UUID
	typ        Type
	assessment assessment.Type
	config     json.RawMessage
}

/*
Constructors
*/

func NewElement(
	lessonID uuid.UUID,
	parentID *uuid.UUID,
	typ Type,
	assessmentType assessment.Type,
	config json.RawMessage,
) (*Element, error) {
	if !typeSet.Has(typ) {
		return nil, erraggregation.Join("element type", enum.ErrInvalidKey, typ)
	}

	if !assessment.IsValidType(assessmentType) {
		return nil, erraggregation.Join("assessment type", enum.ErrInvalidKey, assessmentType)
	}

	if err := validateConfig(config); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()

	if err != nil {
		return nil, err
	}

	return &Element{
		id:         id,
		lessonID:   lessonID,
		parentID:   parentID,
		typ:        typ,
		assessment: assessmentType,
		config:     config,
	}, nil
}

func FromRow(scanner storage.Scanner) (*Element, error) {
	el := &Element{}

	err := scanner.Scan(
		&el.id,
		&el.lessonID,
		&el.parentID,
		&el.typ,
		&el.assessment,
		&el.config,
	)

	if err != nil {
		return nil, err
	}

	return el, nil
}

/*
Domain rules
*/
func validateConfig(config json.RawMessage) error {
	if !json.Valid(config) {
		return errors.New("invalid json")
	}

	return nil
}

/*
Setters
*/

func (e *Element) ChangeParentID(id *uuid.UUID) error {
	if id != nil && e.id == *id {
		return domainerr.NewInvalidFieldValue("ParentID", "cannot be parent of itself")
	}

	e.parentID = id

	return nil
}

func (e *Element) ChangeConfig(config json.RawMessage) error {
	if err := validateConfig(config); err != nil {
		return err
	}

	e.config = config

	return nil
}

/*
Getters
*/

func (e *Element) ID() uuid.UUID { return e.id }

func (e *Element) LessonID() uuid.UUID { return e.lessonID }

func (e *Element) ParentID() *uuid.UUID { return e.parentID }

func (e *Element) Type() Type { return e.typ }

func (e *Element) Assessment() assessment.Type { return e.assessment }

func (e *Element) ConfigRaw() json.RawMessage { return slices.Clone(e.config) }

func (e *Element) ConfigUnmarshal(dst any) error { return json.Unmarshal(e.config, dst) }
