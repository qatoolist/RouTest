package models

import (
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/qatoolist/RouTest/internal/interfaces"
	"gopkg.in/yaml.v3"
)

// Requirements is a map of Requirement objects, indexed by their name.
type Requirements struct {
	sync.RWMutex
	requirements map[string]interfaces.Requirement
}

// NewRequirements returns a new instance of Requirements.
func NewRequirements() *Requirements {
	return &Requirements{
		requirements: make(map[string]interfaces.Requirement),
	}
}

// NewRequirements loads a set of requirements from a YAML file.
func NewRequirementsFromPath(path string) (*Requirements, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var reqMap map[string]Requirement
	err = yaml.Unmarshal(data, &reqMap)
	if err != nil {
		return nil, err
	}

	reqs := NewRequirements()
	for name, req := range reqMap {
		reqs.AddRequirement(name, &req)
	}

	return reqs, nil
}

// AddRequirement adds a new requirement to the Requirements map.
func (r *Requirements) AddRequirement(name string, req interfaces.Requirement) {
	r.Lock()
	defer r.Unlock()
	r.requirements[name] = req
}

// GetRequirement returns a requirement from the Requirements map.
func (r *Requirements) GetRequirement(name string) (interfaces.Requirement, error) {
	r.RLock()
	defer r.RUnlock()
	req, ok := r.requirements[name]
	if !ok {
		return nil, fmt.Errorf("requirement '%s' not found", name)
	}
	return req, nil
}

// RemoveRequirement removes a requirement from the Requirements map.
func (r *Requirements) RemoveRequirement(name string) {
	r.Lock()
	defer r.Unlock()
	delete(r.requirements, name)
}
