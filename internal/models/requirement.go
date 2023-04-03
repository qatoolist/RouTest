package models

import (
	"gopkg.in/yaml.v3"
)

// Requirement represents a single requirement.
type Requirement struct {
	// Summary provides a brief summary of the requirement.
	Summary string `yaml:"summary"`

	// Priority specifies the priority of the requirement.
	Priority string `yaml:"priority"`

	Description string   `yaml:"description"`
	Links       []string `yaml:"links"`
}

// String returns the YAML representation of the requirement.
func (r *Requirement) String() string {
	data, _ := yaml.Marshal(r)
	return string(data)
}
