package models

import (
	"encoding/json"

	"github.com/invopop/jsonschema"
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
