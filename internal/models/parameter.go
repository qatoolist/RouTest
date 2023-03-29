package models

import (
	"errors"
	"net/http"
	"strings"
)

// ParameterType represents the type of a parameter.
type ParameterType struct {
	Key   string
	Value string
}

// Parameter represents a parameter that can be part of an HTTP request.
type Parameter struct {
	QueryParameter  ParameterType
	PathVariable    ParameterType
	HeaderParameter ParameterType
}

// Parameters represents a collection of parameters that can be part of an HTTP request.
type Parameters struct {
	Params []Parameter
}

// NewParameter creates a new Parameter object with the given query parameter, path variable and header.
func NewParameter(paramType string, key string, value string) (*Parameter, error) {
	if key == "" {
		return nil, errors.New("key cannot be empty")
	}
	if value == "" {
		return nil, errors.New("value cannot be empty")
	}
	if paramType != "Query" && paramType != "Path" && paramType != "Header" {
		return nil, errors.New("invalid parameter type")
	}

	param := Parameter{}
	if paramType == "Query" {
		param.QueryParameter = ParameterType{
			Key:   key,
			Value: value,
		}
	} else if paramType == "Header" {
		param.HeaderParameter = ParameterType{
			Key:   key,
			Value: value,
		}
	} else {
		param.PathVariable = ParameterType{
			Key:   key,
			Value: value,
		}
	}

	return &param, nil
}

// AddParameter adds a new parameter to the Parameters object.
func (p *Parameters) AddParameter(param Parameter) {
	p.Params = append(p.Params, param)
}

// GetParameterByKey retrieves the value of a key from the Parameters object for the given ParameterType.
func (p *Parameters) GetParameterByKey(key string, paramType string) (string, error) {
	if key == "" {
		return "", errors.New("key cannot be empty")
	}
	if paramType != "Query" && paramType != "Path" && paramType != "Header" {
		return "", errors.New("invalid parameter type")
	}

	for _, param := range p.Params {
		if paramType == "Query" {
			if param.QueryParameter.Key == key {
				return param.QueryParameter.Value, nil
			}
		} else if paramType == "Header" {
			if param.HeaderParameter.Key == key {
				return param.HeaderParameter.Value, nil
			}
		} else {
			if param.PathVariable.Key == key {
				return param.PathVariable.Value, nil
			}
		}
	}

	return "", errors.New("parameter not found")
}

// AddQueryParamsToRequest adds query parameters to an HTTP request.
func (p *Parameters) AddQueryParamsToRequest(req *http.Request) {
	q := req.URL.Query()
	for _, param := range p.Params {
		if param.QueryParameter.Key != "" {
			q.Add(param.QueryParameter.Key, param.QueryParameter.Value)
		}
	}
	req.URL.RawQuery = q.Encode()
}

// AddPathVarsToRequest adds path variables to an HTTP request.
func (p *Parameters) AddPathVarsToRequest(req *http.Request) {
	for _, param := range p.Params {
		if param.PathVariable.Key != "" {
			req.URL.Path = strings.Replace(req.URL.Path, "{"+param.PathVariable.Key+"}", param.PathVariable.Value, -1)
		}
	}
}

// AddHeaderToRequest adds header parameters to an HTTP request
func (p *Parameters) AddHeaderToRequest(req *http.Request) {
	for _, param := range p.Params {
		if param.HeaderParameter.Key != "" {
			req.Header.Set(param.HeaderParameter.Key, param.HeaderParameter.Value)
		}
	}
}

// ToModelResponse copies all response headers into HeaderParameter
func (p *Parameters) ToModelResponse(resp *http.Response) {
	for key, values := range resp.Header {
		for _, value := range values {
			headerParam, err := NewParameter("Header", key, value)
			if err == nil {
				p.AddParameter(*headerParam)
			}
		}
	}
}
