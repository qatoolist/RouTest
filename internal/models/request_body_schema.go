package models

import (
	"encoding/json"
	"errors"

	"github.com/xeipuuv/gojsonschema"
)

type RequestBodySchema struct {
	SchemaLoader gojsonschema.JSONLoader
}

func NewRequestBodySchema(schema string) (*RequestBodySchema, error) {
	// Load the JSON schema into a JSON loader
	schemaLoader := gojsonschema.NewStringLoader(schema)

	// Validate the schema
	_, err := gojsonschema.NewSchema(schemaLoader)
	if err != nil {
		return nil, err
	}

	// Create a new RequestBodySchema instance
	return &RequestBodySchema{
		SchemaLoader: schemaLoader,
	}, nil
}

func (r *RequestBodySchema) Validate(data interface{}) error {
	// Load the data into a JSON document loader
	documentLoader := gojsonschema.NewGoLoader(data)

	// Validate the data against the schema
	result, err := gojsonschema.Validate(r.SchemaLoader, documentLoader)
	if err != nil {
		return err
	}

	// Check if the data is valid
	if result.Valid() {
		return nil
	}

	// Return a validation error
	var validationErrors []string
	for _, err := range result.Errors() {
		validationErrors = append(validationErrors, err.String())
	}
	return errors.New("validation error: " + validationErrors[0])
}

func (r *RequestBodySchema) MarshalJSON() ([]byte, error) {
	// Marshal the JSON schema to a JSON string
	return json.Marshal(r.SchemaLoader.JsonSource())
}

func (r *RequestBodySchema) UnmarshalJSON(data []byte) error {
	// Load the JSON schema into a JSON loader
	schemaLoader := gojsonschema.NewBytesLoader(data)

	// Create a new RequestBodySchema instance
	r.SchemaLoader = schemaLoader
	return nil
}

/* Example usage -

// Define the JSON schema for the request body
requestBodySchema := `{
    "type": "object",
    "properties": {
        "name": {"type": "string"},
        "age": {"type": "number"},
        "email": {"type": "string", "format": "email"}
    },
    "required": ["name", "age"]
}`

// Create a new RequestBodySchema instance
schema, err := NewRequestBodySchema(requestBodySchema)
if err != nil {
    log.Fatal(err)
}

// Validate a request body against the schema
requestBody := map[string]interface{}{
    "name": "John",
    "age": 30,
    "email": "john@example.com",
}
err = schema.Validate(requestBody)
if err != nil {
    log.Fatal(err)
}

// Marshal the schema to a JSON string
schemaJSON, err := json.Marshal(schema)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(schemaJSON))

*/
