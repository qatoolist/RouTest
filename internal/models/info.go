package models

// Info represents the information about an API endpoint.
type Info struct {
	// Name is the name of the API endpoint.
	Name string

	// Description is a brief description of the API endpoint.
	Description string

	// Path is the URL path of the API endpoint.
	Path string

	// Method is the HTTP method used by the API endpoint.
	Method Method

	// Params contains the parameters of the API endpoint.
	Params Parameters

	// RequestBodySchema is the JSON schema for the request body of the API endpoint.
	RequestBodySchema RequestBodySchema

	// ResponseBodySchema is the JSON schema for the response body of the API endpoint.
	ResponseBodySchema ResponseBodySchema
}
