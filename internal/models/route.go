package models

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Route represents a custom HTTP request with additional fields.
type Route struct {
	// Info contains the request information.
	Info Info

	// Meta contains the metadata about the route.
	Meta Meta

	// Scenario contains the scenario that should be executed for the route.
	Scenario Scenario

	// Body is the request body.
	Body []byte

	// Response is the channel for the HTTP response.
	Response <-chan *http.Response
}

// NewRoute creates a new instance of the Route struct with default values for its fields.
func NewRoute() *Route {
	return &Route{
		Info:     Info{},
		Meta:     Meta{},
		Scenario: Scenario{},
		Body:     []byte{},
		Response: nil,
	}
}

// SetReqBodySchema sets the request body schema for the route.
func (r *Route) SetReqBodySchema(schema string) error {
	reqBodySchema, err := NewRequestBodySchema(schema)
	if err != nil {
		return err
	}
	r.Info.RequestBodySchema = *reqBodySchema
	return nil
}

// SetResBodySchema sets the response body schema for the route.
func (r *Route) SetResBodySchema(schema string) error {
	resBodySchema, err := NewResponseBodySchema(schema)
	if err != nil {
		return err
	}
	r.Info.ResponseBodySchema = *resBodySchema
	return nil
}

// ValidateReqBody validates the request body against the request body schema.
func (r *Route) ValidateReqBody(body interface{}) error {
	if r.Info.RequestBodySchema == (RequestBodySchema{}) {
		return nil
	}
	return r.Info.RequestBodySchema.Validate(body)
}

// ValidateResBody validates the response body against the response body schema.
func (r *Route) ValidateResBody(body interface{}) error {
	if r.Info.ResponseBodySchema == (ResponseBodySchema{}) {
		return nil
	}
	return r.Info.ResponseBodySchema.Validate(body)
}

// Send sends the HTTP request and returns the HTTP response.
func (r *Route) Send() (*http.Response, error) {
	client := http.DefaultClient
	req, err := http.NewRequest(r.Info.Method.String(), r.Info.Path, bytes.NewReader(r.Body))
	if err != nil {
		return nil, err
	}

	r.Info.Params.AddHeaderToRequest(req)

	if r.Info.RequestBodySchema != (RequestBodySchema{}) {
		err := r.ValidateReqBody(req.Body)
		if err != nil {
			return nil, err
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if r.Info.ResponseBodySchema != (ResponseBodySchema{}) {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		err = r.ValidateResBody(body)
		if err != nil {
			return nil, err
		}
		resp.Body = ioutil.NopCloser(bytes.NewReader(body))
	}
	return resp, nil
}
