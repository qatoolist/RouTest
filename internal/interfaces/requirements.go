package interfaces

// Requirements is a map of Requirement objects, indexed by their name.
type Requirements interface {
	AddRequirement(name string, req Requirement)
	GetRequirement(name string) (Requirement, error)
	RemoveRequirement(name string)
}
