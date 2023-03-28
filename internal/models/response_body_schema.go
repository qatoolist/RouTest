package models

import (
	"encoding/json"

	"github.com/alecthomas/jsonschema"
)

type ResponseBodySchema struct {
	jsonschema.Schema
}

func (r *ResponseBodySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.Schema)
}

func (r *ResponseBodySchema) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Schema)
}

func (r *ResponseBodySchema) Validate(data interface{}) error {
	return r.Schema.Validate(data)
}

/*
example -

schema := &models.ResponseBodySchema{}
err := json.Unmarshal([]byte(`
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "type": "object",
  "properties": {
    "name": {
      "type": "string"
    },
    "age": {
      "type": "integer"
    }
  },
  "required": ["name"]
}
`), schema)
if err != nil {
  panic(err)
}

*/
