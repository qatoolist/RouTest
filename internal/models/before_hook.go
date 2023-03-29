package models

// BeforeHook is a function that takes a Route object as argument and returns a modified Route object with or without error.
type BeforeHook func(*Route) (*Route, error)

// Execute method executes the BeforeHook function with the given Route as argument.
// It returns the modified Route object and any error occurred during the execution of BeforeHook.
func (bh BeforeHook) Execute(route *Route) (*Route, error) {
	return bh(route)
}
