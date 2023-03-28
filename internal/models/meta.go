package models

import (
	"errors"

	"gopkg.in/yaml.v3"
)

type AutomationStatus string

const (
	Automated    AutomationStatus = "automated"
	NotAutomated AutomationStatus = "not_automated"
	ManualOnly   AutomationStatus = "manual_only"
)

type Importance string

const (
	Critical Importance = "critical"
	High     Importance = "high"
	Medium   Importance = "medium"
	Low      Importance = "low"
)

type Meta struct {
	Assignee             string           `json:"assignee"`
	AutomationStatus     AutomationStatus `json:"automation_status"`
	Component            string           `json:"component"`
	Importance           Importance       `json:"importance"`
	Requirements         string           `json:"requirements"`
	RequirementsOverride string           `json:"requirements_override"`
	Setup                string           `json:"setup"`
	TestSteps            string           `json:"test_steps"`
	ExpectedResults      string           `json:"expected_results"`
	Negative             bool             `json:"negative"`
	Type                 string           `json:"type"`
}

// NewMetaFromString creates a new Meta struct from a YAML string.
func (m *Meta) NewMetaFromString(s string) error {
	if s == "" {
		return errors.New("empty input string")
	}

	err := yaml.Unmarshal([]byte(s), m)
	if err != nil {
		return err
	}

	// Check for invalid enum values
	if m.AutomationStatus != Automated && m.AutomationStatus != NotAutomated && m.AutomationStatus != ManualOnly {
		return errors.New("invalid automation_status value")
	}
	if m.Importance != Critical && m.Importance != High && m.Importance != Medium && m.Importance != Low {
		return errors.New("invalid importance value")
	}

	return nil
}
