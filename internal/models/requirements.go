package models

// Requirements is a map of Requirement objects, indexed by their name.
type Requirements map[string]Requirement

// AddRequirement adds a new requirement to the Requirements map.
func (r *Requirements) AddRequirement(name string, req *Requirement) {
	(*r)[name] = *req
}

// GetRequirement returns a requirement from the Requirements map.
func (r *Requirements) GetRequirement(name string) *Requirement {
	if req, ok := (*r)[name]; ok {
		return &req
	}
	return nil
}

// RemoveRequirement removes a requirement from the Requirements map.
func (r *Requirements) RemoveRequirement(name string) {
	delete(*r, name)
}
