package models

import (
	"encoding/json"
	"errors"

	"github.com/xeipuuv/gojsonschema"
)

// ResponseBodySchema represents the schema for validating a response body.
type ResponseBodySchema struct {
	SchemaLoader gojsonschema.JSONLoader // schema loader for JSON schema
}

// NewResponseBodySchema creates a new instance of ResponseBodySchema.
func NewResponseBodySchema(schema string) (*ResponseBodySchema, error) {
	// Load the JSON schema into a JSON loader
	schemaLoader := gojsonschema.NewStringLoader(schema)

	// Validate the schema
	_, err := gojsonschema.NewSchema(schemaLoader)
	if err != nil {
		return nil, err
	}

	// Create a new ResponseBodySchema instance
	return &ResponseBodySchema{
		SchemaLoader: schemaLoader,
	}, nil
}

// Validate validates the given data against the schema.
func (r *ResponseBodySchema) Validate(data interface{}) error {
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

// MarshalJSON marshals the JSON schema to a JSON string.
func (r *ResponseBodySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.SchemaLoader.JsonSource())
}

// UnmarshalJSON unmarshals the JSON schema from a JSON string.
func (r *ResponseBodySchema) UnmarshalJSON(data []byte) error {
	// Load the JSON schema into a JSON loader
	schemaLoader := gojsonschema.NewBytesLoader(data)

	// Create a new ResponseBodySchema instance
	r.SchemaLoader = schemaLoader
	return nil
}

/* Example usage -

// Define the JSON schema for the response body
responseBodySchema := `{
    "type": "object",
    "properties": {
        "name": {"type": "string"},
        "age": {"type": "number"},
        "email": {"type": "string", "format": "email"}
    },
    "required": ["name", "age"]
}`

// Create a new ResponseBodySchema instance
schema, err := NewResponseBodySchema(responseBodySchema)
if err != nil {
    log.Fatal(err)
}

// Validate a response body against the schema
responseBody := map[string]interface{}{
    "name": "John",
    "age": 30,
    "email": "john@example.com",
}
err = schema.Validate(responseBody)
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
