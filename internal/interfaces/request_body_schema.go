package interfaces

type RequestBodySchema interface {
	Validate(interface{}) error
}
