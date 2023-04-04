package models

import (
	"errors"

	"github.com/qatoolist/RouTest/internal/interfaces"
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
	ParentMeta interfaces.Meta `json:"parent_meta" yaml:"parent_meta"`

	// The person assigned to the task
	Assignee string `json:"assignee" yaml:"assignee"`

	// Whether the test is automated, not automated, or manual-only
	AutomationStatus AutomationStatus `json:"automation_status" yaml:"automation_status"`

	// The component being tested
	Component string `json:"component" yaml:"component"`

	// The importance of the task (e.g. critical, high, medium, low)
	Importance Importance `json:"importance" yaml:"importance"`

	// The requirements being tested
	Requirements string `json:"requirements" yaml:"requirements"`

	// Whether the requirements for this task have been overridden
	RequirementsOverride string `json:"requirements_override" yaml:"requirements_override"`

	// The setup required for the task
	Setup string `json:"setup" yaml:"setup"`

	// The steps to perform the test
	TestSteps string `json:"test_steps" yaml:"test_steps"`

	// The expected results of the test
	ExpectedResults string `json:"expected_results" yaml:"expected_results"`

	// Whether the test is negative
	Negative bool `json:"negative" yaml:"negative"`

	// The type of the test (e.g. smoke, regression, acceptance)
	Type string `json:"type" yaml:"type"`

	// The tags associated with the task
	Tags string `json:"tags" yaml:"tags"`
}

func NewMetaFromString(s string) (interfaces.Meta, error) {
	if s == "" {
		return nil, errors.New("empty input string")
	}

	var m Meta

	err := yaml.Unmarshal([]byte(s), &m)
	if err != nil {
		return nil, err
	}

	// Check for invalid enum values
	if m.AutomationStatus != Automated && m.AutomationStatus != NotAutomated && m.AutomationStatus != ManualOnly {
		return nil, errors.New("invalid automation_status value")
	}
	if m.Importance != Critical && m.Importance != High && m.Importance != Medium && m.Importance != Low {
		return nil, errors.New("invalid importance value")
	}

	return &m, nil
}

func (m *Meta) Copy() interfaces.Meta {
	copy := &Meta{
		ParentMeta:           m.ParentMeta,
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

	return copy
}

func (m *Meta) OverrideMeta(override interfaces.Meta) {
	if override == nil {
		return
	}

	om, ok := override.(*Meta)
	if !ok {
		return
	}

	m.Assignee = om.Assignee
	m.AutomationStatus = om.AutomationStatus
	m.Component = om.Component
	m.Importance = om.Importance
	m.Requirements = om.Requirements
	m.RequirementsOverride = om.RequirementsOverride
	m.Setup = om.Setup
	m.TestSteps = om.TestSteps
	m.ExpectedResults = om.ExpectedResults
	m.Negative = om.Negative
	m.Type = om.Type

	if om.ParentMeta != nil {
		m.ParentMeta = om.ParentMeta.Copy()
	}

	// Append or override the tags
	if om.Tags != "" {
		if m.Tags != "" {
			m.Tags = m.Tags + "," + om.Tags
		} else {
			m.Tags = om.Tags
		}
	}
}

func (m *Meta) GetParentMeta() interfaces.Meta {
	return m.ParentMeta
}

func (m *Meta) GetTags() string {
	return m.Tags
}
