// Package models provides types for defining HTTP scenarios.
package models

import (
	"encoding/json"
)

// Scenario defines an HTTP scenario.
type Scenario struct {
	// Meta provides metadata about the scenario.
	Meta *Meta

	// RequestBodySchema specifies the schema for the request body.
	RequestBodySchema RequestBodySchema

	// ResponseBodySchema specifies the schema for the response body.
	ResponseBodySchema ResponseBodySchema

	// BeforeHook is a function that will be executed before the request is sent.
	BeforeHook BeforeHook

	// AfterHook is a function that will be executed after the response is received.
	AfterHook AfterHook

	// Response represents the HTTP Response recieved fter sending the request for this scenario
	Response *Response
}

// NewScenario creates a new Scenario object with the given metadata and default values for other fields.
func NewScenario(meta *Meta) *Scenario {
	return &Scenario{
		Meta:               meta.Copy(),
		RequestBodySchema:  RequestBodySchema{},
		ResponseBodySchema: ResponseBodySchema{},
		BeforeHook:         nil,
		AfterHook:          nil,
		Response:           nil,
	}
}

// ExecuteBeforeHook executes the before hook for the scenario.
func (s *Scenario) ExecuteBeforeHook(route *Route) (*Route, error) {
	if s.BeforeHook != nil {
		return s.BeforeHook.Execute(route)
	}
	return nil, nil
}

// ExecuteAfterHook executes the after hook for the scenario.
func (s *Scenario) ExecuteAfterHook(resp *Response) (*Response, error) {
	if s.AfterHook != nil {
		return s.AfterHook.Execute(resp)
	}
	return nil, nil
}

// ValidateResponseBody validates the response body against the schema.
func (s *Scenario) ValidateResponseBody(body []byte, schema ResponseBodySchema) error {
	if schema != (ResponseBodySchema{}) {
		var data interface{}
		if err := json.Unmarshal(body, &data); err != nil {
			return err
		}
		if err := schema.Validate(data); err != nil {
			return err
		}
	}
	return nil
}
