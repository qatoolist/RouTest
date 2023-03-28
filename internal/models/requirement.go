package models

type Requirement struct {
	Summary  string   `json:"summary"`
	Priority string   `json:"priority"`
	Links    []string `json:"links"`
}

// GetPriority returns the priority of the requirement.
func (r *Requirement) GetPriority() string {
	return r.Priority
}

// SetPriority sets the priority of the requirement.
func (r *Requirement) SetPriority(priority string) {
	r.Priority = priority
}

// AddLink adds a link to the requirement.
func (r *Requirement) AddLink(link string) {
	r.Links = append(r.Links, link)
}

// GetSummary returns the summary of the requirement.
func (r *Requirement) GetSummary() string {
	return r.Summary
}

// SetSummary sets the summary of the requirement.
func (r *Requirement) SetSummary(summary string) {
	r.Summary = summary
}