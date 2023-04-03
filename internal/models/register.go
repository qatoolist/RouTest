package models

import (
	"fmt"

	"github.com/qatoolist/RouTest/internal/interfaces"
)

// Register is a collection of Responses and Parameters.
type Register struct {
	Responses  interfaces.ResponseMap
	Parameters interfaces.ParameterMap
}

//invalid composite literal type interfaces.Register

func NewRegister() interfaces.Register {
	return &Register{
		Responses:  make(interfaces.ResponseMap),
		Parameters: make(interfaces.ParameterMap),
	}
}

// AddResponse adds a new response to the register with the given name.
func (r *Register) AddResponse(name string, resp *interfaces.Response) {
	r.Responses[name] = resp
}

// GetResponse retrieves the response with the given name.
func (r *Register) GetResponse(name string) (*interfaces.Response, error) {
	resp, ok := r.Responses[name]
	if !ok {
		return nil, fmt.Errorf("response '%s' not found", name)
	}
	return resp, nil
}

// AddParameter adds a new parameter to the register with the given name.
func (r *Register) AddParameter(name string, param *interfaces.Parameter) {
	r.Parameters[name] = param
}

// GetParameter retrieves the parameter with the given name.
func (r *Register) GetParameter(name string) (*interfaces.Parameter, error) {
	param, ok := r.Parameters[name]
	if !ok {
		return nil, fmt.Errorf("parameter '%s' not found", name)
	}
	return param, nil
}
