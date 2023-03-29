package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Response represents the HTTP response and its attributes.
type Response struct {
	// StatusCode is the HTTP status code of the response.
	StatusCode int

	// Status is the HTTP status message of the response.
	Status string

	// Parameters is a collection of parameters associated with the response.
	Parameters Parameters

	// Body is the body of the response.
	Body []byte
}

// NewResponse creates a new instance of the Response struct with default values for its fields.
func NewResponse(rbs ResponseBodySchema) (*Response, error) {
	resp := &Response{
		StatusCode: -1,
		Status:     "",
		Body:       nil,
	}
	return resp, nil
}

// HandleResponse handles the HTTP response and returns the Response object.
func HandleResponse(httpResp *http.Response) (*Response, error) {
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	resp := &Response{
		StatusCode: httpResp.StatusCode,
		Status:     httpResp.Status,
		Body:       body,
	}

	resp.Parameters = Parameters{}
	resp.Parameters.AddParameter(Parameter{
		HeaderParameter: ParameterType{
			Key:   "Content-Type",
			Value: httpResp.Header.Get("Content-Type"),
		},
	})

	resp.Parameters.ToModelResponse(httpResp)

	return resp, nil
}

// String returns the response body as a string.
func (r *Response) String() string {
	return string(r.Body)
}

// Bytes returns the response body as a byte slice.
func (r *Response) Bytes() []byte {
	return r.Body
}

// HeaderValue returns the value of the specified header.
func (r *Response) HeaderValue(key string) (string, error) {
	return r.Parameters.GetParameterByKey(key, "Header")
}

// ContentType returns the Content-Type header value.
func (r *Response) ContentType() (string, error) {
	return r.HeaderValue("Content-Type")
}

// IsSuccess returns true if the response status code indicates success.
func (r *Response) IsSuccess() bool {
	return r.StatusCode >= 200 && r.StatusCode < 300
}

// SetHeaderValue sets the value of the specified header.
func (r *Response) SetHeaderValue(key, value string) {
	headerParam, err := NewParameter("Header", key, value)
	if err == nil {
		r.Parameters.AddParameter(*headerParam)
	}
}

// ValidateBody validates the response body against the schema.
func (r *Response) ValidateBody(schema *ResponseBodySchema) error {
	var data interface{}

	if err := json.Unmarshal(r.Body, &data); err != nil {
		return err
	}

	if err := schema.Validate(data); err != nil {
		return err
	}

	return nil
}
