package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/qatoolist/RouTest/internal/interfaces"
)

// Response represents the HTTP response and its attributes.
type Response struct {
	// StatusCode is the HTTP status code of the response.
	StatusCode int

	// Status is the HTTP status message of the response.
	Status string

	// ResponseParametersRegistry is a collection of parameters associated with the response.
	ResponseParametersRegistry interfaces.ParametersRegistry

	// Body is the body of the response.
	Body []byte
}

// NewResponse creates a new instance of the Response struct with default values for its fields.
func NewResponse() (interfaces.Response, error) {
	resp := &Response{
		StatusCode:                 -1,
		Status:                     "",
		ResponseParametersRegistry: NewParameterRegistry(),
		Body:                       nil,
	}
	return interfaces.Response(resp), nil
}

// HandleResponse handles the HTTP response and returns the Response object.
func HandleResponse(httpResp *http.Response) (interfaces.Response, error) {
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	resp := &Response{
		StatusCode: httpResp.StatusCode,
		Status:     httpResp.Status,
		Body:       body,
	}

	resp.ResponseParametersRegistry = NewParameterRegistry()
	resp.ResponseParametersRegistry.RegisterHeader("Content-Type", "Content-Type")

	resp.ResponseParametersRegistry.ImportFromHTTPResponse(httpResp)

	return interfaces.Response(resp), nil
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
	return r.ResponseParametersRegistry.GetParameterByKey(key, "Header")
}

// ContentType returns the Content-Type header value.
func (r *Response) ContentType() (string, error) {
	return r.HeaderValue("Content-Type")
}

// IsSuccess returns true if the response status code indicates success.
func (r *Response) IsSuccess() bool {
	return r.StatusCode >= 200 && r.StatusCode < 300
}

// ValidateBody validates the response body against the schema.
func (r *Response) ValidateBody(schema interfaces.ResponseBodySchema) error {
	var data interface{}

	if err := json.Unmarshal(r.Body, &data); err != nil {
		return err
	}

	if err := schema.Validate(data); err != nil {
		return err
	}

	return nil
}
