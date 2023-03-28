package models

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type Response struct {
	StatusCode int
	Status     string
	Proto      string
	ProtoMajor int
	ProtoMinor int
	Header     Header
	Body       []byte
}

func ToModelHeader(h http.Header) Header {
	header := Header{}
	for key, values := range h {
		header[key] = values
	}
	return header
}

func NewResponse(httpResp *http.Response) (*Response, error) {
	bodyBytes, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	return &Response{
		StatusCode: httpResp.StatusCode,
		Status:     httpResp.Status,
		Proto:      httpResp.Proto,
		ProtoMajor: httpResp.ProtoMajor,
		ProtoMinor: httpResp.ProtoMinor,
		Header:     ToModelHeader(httpResp.Header),
		Body:       bodyBytes,
	}, nil
}

// String returns the response as a string.
func (r *Response) String() string {
	return string(r.Body)
}

// JSON parses the response body as a JSON object and returns it.
func (r *Response) JSON(v interface{}) (interface{}, error) {
	err := json.Unmarshal(r.Body, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// XML parses the response as an XML object and returns it.
func (r *Response) XML(v interface{}) (interface{}, error) {
	err := xml.Unmarshal(r.Body, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Bytes returns the response as a byte slice.
func (r *Response) Bytes() []byte {
	return r.Body
}

// HeaderValue returns the value of the specified header.
func (r *Response) HeaderValue(key string) string {
	return r.Header.Get(key)
}

// SetHeaderValue sets the value of the specified header.
func (r *Response) SetHeaderValue(key, value string) {
	r.Header.Set(key, value)
}

// AddHeaderValue adds a value to the specified header.
func (r *Response) AddHeaderValue(key, value string) {
	r.Header.Add(key, value)
}

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
