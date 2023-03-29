package models

import "fmt"

// ResponseMap maps response names to their respective Response structs.
type ResponseMap map[string]*Response

// ParameterMap maps parameter names to their respective Parameter structs.
type ParameterMap map[string]*Parameter

// Register is a collection of Responses and Parameters.
type Register struct {
	Responses  ResponseMap
	Parameters ParameterMap
}

// AddResponse adds a new response to the register with the given name.
func (r *Register) AddResponse(name string, resp *Response) {
	r.Responses[name] = resp
}

// GetResponse retrieves the response with the given name.
func (r *Register) GetResponse(name string) (*Response, error) {
	resp, ok := r.Responses[name]
	if !ok {
		return nil, fmt.Errorf("response '%s' not found", name)
	}
	return resp, nil
}

// AddParameter adds a new parameter to the register with the given name.
func (r *Register) AddParameter(name string, param *Parameter) {
	r.Parameters[name] = param
}

// GetParameter retrieves the parameter with the given name.
func (r *Register) GetParameter(name string) (*Parameter, error) {
	param, ok := r.Parameters[name]
	if !ok {
		return nil, fmt.Errorf("parameter '%s' not found", name)
	}
	return param, nil
}
