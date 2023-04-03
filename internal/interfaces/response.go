package interfaces

type Response interface {
	String() string
	Bytes() []byte
	HeaderValue(key string) (string, error)
	ContentType() (string, error)
	IsSuccess() bool
	ValidateBody(schema ResponseBodySchema) error
}
