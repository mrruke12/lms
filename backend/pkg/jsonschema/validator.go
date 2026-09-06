package jsonschema

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/mrruke12/lms/pkg/enum"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

// JSON Schema validator
type JSONValidator[T ~string] struct {
	compiler *jsonschema.Compiler
	set      *enum.Set[T]
	schemas  map[T]*jsonschema.Schema
}

func NewJSONValidator[T ~string](set *enum.Set[T]) *JSONValidator[T] {
	return &JSONValidator[T]{
		compiler: jsonschema.NewCompiler(),
		set:      set,
		schemas:  make(map[T]*jsonschema.Schema),
	}
}

// Add a new key-schema pair
func (v *JSONValidator[T]) RegisterSchema(key T, shemaString string) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Could not register schema for %s: %w", key, err)
		}
	}()

	if !v.set.Has(key) {
		return enum.ErrInvalidKey
	}

	parsedSchema, err := jsonschema.UnmarshalJSON(
		bytes.NewReader(
			[]byte(shemaString),
		),
	)

	if err != nil {
		return err
	}

	name := string(key) + ".json"

	if err := v.compiler.AddResource(name, parsedSchema); err != nil {
		return err
	}

	schema, err := v.compiler.Compile(name)

	if err != nil {
		return err
	}

	v.schemas[key] = schema

	return nil
}

// Validate the given raw json by key schema
func (v *JSONValidator[T]) Validate(key T, raw json.RawMessage) error {
	if !v.set.Has(key) {
		return enum.ErrInvalidKey
	}

	schema, exists := v.schemas[key]

	if !exists {
		return ErrSchemaMissing
	}

	inputData, err := jsonschema.UnmarshalJSON(bytes.NewReader(raw))

	if err != nil {
		return err
	}

	if err := schema.Validate(inputData); err != nil {
		return err
	}

	return nil
}
