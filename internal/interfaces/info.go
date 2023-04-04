package interfaces

type Info interface {
	GetName() string

	GetDescription() string

	GetPath() string

	GetMethod() Method

	GetRequestBodySchema() RequestBodySchema

	GetResponseBodySchema() ResponseBodySchema

	SetName(string)

	SetDescription(string)

	SetPath(string)

	SetMethod(Method)

	SetRequestBodySchema(RequestBodySchema)

	SetResponseBodySchema(ResponseBodySchema)
}
