package interfaces

type ResponseBodySchema interface {
	Validate(interface{}) error
}
