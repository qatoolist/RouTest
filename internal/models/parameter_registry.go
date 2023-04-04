package models

import (
	"errors"
	"net/http"
	"strings"
	"sync"

	"github.com/qatoolist/RouTest/internal/interfaces"
)

// Parameter represents a parameter that can be part of an HTTP request.
type ParameterRegistry struct {
	queryParameters  []interfaces.Parameter
	pathVariables    []interfaces.Parameter
	headerParameters []interfaces.Parameter
	mu               sync.RWMutex
}

func NewParameterRegistry() interfaces.ParametersRegistry {
	return &ParameterRegistry{
		queryParameters:  make([]interfaces.Parameter, 0),
		pathVariables:    make([]interfaces.Parameter, 0),
		headerParameters: make([]interfaces.Parameter, 0),
	}
}

// ImportFromHTTPResponse imports parameters from an HTTP response.
func (pr *ParameterRegistry) ImportFromHTTPResponse(httpResp *http.Response) error {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	// Query parameters
	qp := httpResp.Request.URL.Query()
	for key, values := range qp {
		for _, value := range values {
			pr.RegisterQueryParameter(key, value)
		}
	}

	// Header parameters
	for key, values := range httpResp.Header {
		for _, value := range values {
			pr.RegisterHeader(key, value)
		}
	}

	return nil
}

// RegisterQueryParameter adds a query parameter to the registry.
func (pr *ParameterRegistry) RegisterQueryParameter(key string, value string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}
	pr.mu.Lock()
	defer pr.mu.Unlock()
	pr.queryParameters = append(pr.queryParameters, NewParameter(key, value))
	return nil
}

// RegisterPathVariable adds a path variable to the registry.
func (pr *ParameterRegistry) RegisterPathVariable(key string, value string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}
	pr.mu.Lock()
	defer pr.mu.Unlock()
	pr.pathVariables = append(pr.pathVariables, NewParameter(key, value))

	return nil
}

// RegisterHeader adds a header parameter to the registry.
func (pr *ParameterRegistry) RegisterHeader(key string, value string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}
	pr.mu.Lock()
	defer pr.mu.Unlock()
	pr.headerParameters = append(pr.headerParameters, NewParameter(key, value))
	return nil
}

// GetQueryParameters returns all the registered query parameters.
func (pr *ParameterRegistry) GetQueryParameters() []interfaces.Parameter {
	pr.mu.RLock()
	defer pr.mu.RUnlock()
	return pr.queryParameters
}

// GetPathVariables returns all the registered path variables.
func (pr *ParameterRegistry) GetPathVariables() []interfaces.Parameter {
	pr.mu.RLock()
	defer pr.mu.RUnlock()
	return pr.pathVariables
}

// GetParameters returns all the registered parameters.
func (pr *ParameterRegistry) GetHeaders() []interfaces.Parameter {
	pr.mu.RLock()
	defer pr.mu.RUnlock()
	return pr.headerParameters
}

// GetParameterByKey returns the value of the specified parameter key and type.
func (pr *ParameterRegistry) GetParameterByKey(key string, paramType string) (string, error) {
	pr.mu.RLock()
	defer pr.mu.RUnlock()

	// lookup in Query parameters
	if paramType == "Query" && key != "" {
		for _, param := range pr.queryParameters {
			if param.Key() == key {
				return param.Value(), nil
			}
		}
	}

	// lookup in Path Variables
	if paramType == "Path" && key != "" {
		for _, param := range pr.pathVariables {
			if param.Key() == key {
				return param.Value(), nil
			}
		}
	}

	// lookup in Path Variables
	if paramType == "Header" && key != "" {
		for _, param := range pr.headerParameters {
			if param.Key() == key {
				return param.Value(), nil
			}
		}
	}

	return "", errors.New("parameter not found")
}

// ExportToRequest exports the parameters to an HTTP request.
func (pr *ParameterRegistry) ExportToRequest(req *http.Request) (*http.Request, error) {
	pr.mu.RLock()
	defer pr.mu.RUnlock()

	q := req.URL.Query()

	for _, param := range pr.queryParameters {
		if param.Key() != "" {
			q.Add(param.Key(), param.Value())
		}
	}
	for _, param := range pr.pathVariables {
		if param.Key() != "" {
			pathVarKey := "{" + param.Key() + "}"
			req.URL.Path = strings.Replace(req.URL.Path, pathVarKey, param.Value(), -1)
		}
	}
	for _, param := range pr.headerParameters {
		if param.Key() != "" {
			req.Header.Set(param.Key(), param.Value())
		}
	}
	req.URL.RawQuery = q.Encode()
	return req, nil
}
