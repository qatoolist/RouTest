package models

import (
	"github.com/qatoolist/RouTest/internal/interfaces"
)

// ParameterType represents the type of a parameter.
type Parameter struct {
	key   string
	value string
}

func (p *Parameter) Key() string {
	return p.key
}

func (p *Parameter) Value() string {
	return p.value
}

func NewParameter(key string, value string) interfaces.Parameter {
	return &Parameter{
		key:   key,
		value: value,
	}
}
