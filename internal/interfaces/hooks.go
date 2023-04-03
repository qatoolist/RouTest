package interfaces

// BeforeHook is a function that takes a Route object as argument and returns a modified Route object with or without error.
type BeforeHook func(Route) (Route, error)

// AfterHook is a function type that takes a Response pointer as argument and returns a Response pointer and an error.
type AfterHook func(Response) (Response, error)
