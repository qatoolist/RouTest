package models

// AfterHook is a function type that takes a Response pointer as argument and returns a Response pointer and an error.
type AfterHook func(*Response) (*Response, error)

// Execute executes the AfterHook with the given Response as argument.
func (ah AfterHook) Execute(resp *Response) (*Response, error) {
	return ah(resp)
}
