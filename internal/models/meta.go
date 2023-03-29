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

	// The reference to the Meta of Parent container i.e. Route or Application
	ParentMeta *Meta

	// The person assigned to the task
	Assignee string `json:"assignee"`

	// Whether the test is automated, not automated, or manual-only
	AutomationStatus AutomationStatus `json:"automation_status"`

	// The component being tested
	Component string `json:"component"`

	// The importance of the task (e.g. critical, high, medium, low)
	Importance Importance `json:"importance"`

	// The requirements being tested
	Requirements string `json:"requirements"`

	// Whether the requirements for this task have been overridden
	RequirementsOverride string `json:"requirements_override"`

	// The setup required for the task
	Setup string `json:"setup"`

	// The steps to perform the test
	TestSteps string `json:"test_steps"`

	// The expected results of the test
	ExpectedResults string `json:"expected_results"`

	// Whether the test is negative
	Negative bool `json:"negative"`

	// The type of the test (e.g. smoke, regression, acceptance)
	Type string `json:"type"`

	// The tags associated with the task
	Tags string `json:"tags"`
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

func (m *Meta) Copy() *Meta {
	return &Meta{
		ParentMeta:           m,
		Assignee:             m.Assignee,
		AutomationStatus:     m.AutomationStatus,
		Component:            m.Component,
		Importance:           m.Importance,
		Requirements:         m.Requirements,
		RequirementsOverride: m.RequirementsOverride,
		Setup:                m.Setup,
		TestSteps:            m.TestSteps,
		ExpectedResults:      m.ExpectedResults,
		Negative:             m.Negative,
		Type:                 m.Type,
		Tags:                 m.Tags,
	}
}
