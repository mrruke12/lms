package element

import (
	"bytes"
	"encoding/json"

	"github.com/mrruke12/lms/internal/apperr"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

// Base schema for any type of element
const elementSchemaString = `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "type": "object",
  "properties": {
    "type": {
      "type": "string"
    },
    "content": {
      "type": "array",
      "minItems": 0,
      "maxItems": 100,
      "prefixItems": [
        {
          "type": "string"
        },
        {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "type": {
                "type": "string",
                "enum": ["text", "medialink"]
              },
              "value": {
                "type": "string"
              }
            },
            "required": ["type", "value"],
            "additionalProperties": false
          }
        }
      ],
      "additionalItems": false
    },
	"assessment": {
		"type": "object"
	},
    "styles": {
      "type": "array",
      "items": {
        "type": "array",
        "minItems": 2,
        "maxItems": 2,
        "prefixItems": [
          { "type": "string" },
          { "type": "string" }
        ],
        "additionalItems": false
      }
    }
  },
  "required": ["type", "content", "styles"],
  "additionalProperties": false
}`

// Additional schema to validate assessable elements
const assessmentSchemaString = `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "type": "object",
  "properties": {
    "assessment": {
      "type": "object",
      "additionalProperties": true
    }
  },
  "required": ["assessment"],
  "additionalProperties": true
}`

var compiler = jsonschema.NewCompiler()

// Compiles schema. Panics on error
func compileSchema(schemaString string, name string) *jsonschema.Schema {
	parsedSchema, err := jsonschema.UnmarshalJSON(
		bytes.NewReader(
			[]byte(schemaString),
		),
	)

	if err != nil {
		panic(err)
	}

	if err := compiler.AddResource(name+".json", parsedSchema); err != nil {
		panic(err)
	}

	schema, err := compiler.Compile(name + ".json")

	if err != nil {
		panic(err)
	}

	return schema
}

// Validate raw json
func validateJSON(schema *jsonschema.Schema, raw json.RawMessage) error {
	inputData, err := jsonschema.UnmarshalJSON(bytes.NewReader(raw))

	if err != nil {
		return apperr.InvalidJSONSchema(err.Error())
	}

	if err := schema.Validate(inputData); err != nil {
		return apperr.InvalidJSONSchema(err.Error())
	}

	return nil
}

// Reusable schema to validate Base Element
var elementSchema = compileSchema(elementSchemaString, "base_element")

// Reusable schema to validate Assessment Element
var assessmentSchema = compileSchema(assessmentSchemaString, "assessment_element")
