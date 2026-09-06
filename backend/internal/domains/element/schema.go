package element

import "encoding/json"

// Base schema for any type of element
const elementSchemaString = `
{
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
const assessmentSchemaString = `
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "type": "object",
  "properties": {
    "assessment": {
      "type": "object",
	  "properties": {
	  	"cap": {
			"type": "integer"
		}
	  },
	  "required": ["cap"],
      "additionalProperties": true
    }
  },
  "required": ["assessment"],
  "additionalProperties": true
}`

type SchemaValidator interface {
	Validate(typ string, config json.RawMessage) error
}
