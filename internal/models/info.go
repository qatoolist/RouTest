package models

import (
	"github.com/qatoolist/RouTest/internal/interfaces"
	"gopkg.in/yaml.v3"
)

// Info represents the information about an API endpoint.
type Info struct {
	// Name is the name of the API endpoint.
	name string

	// Description is a brief description of the API endpoint.
	description string

	// Path is the URL path of the API endpoint.
	path string

	// Method is the HTTP method used by the API endpoint.
	method interfaces.Method

	// RequestBodySchema is the JSON schema for the request body of the API endpoint.
	requestBodySchema interfaces.RequestBodySchema

	// ResponseBodySchema is the JSON schema for the response body of the API endpoint.
	responseBodySchema interfaces.ResponseBodySchema
}

func NewInfo(infoStr string) interfaces.Info {
	var info Info
	err := yaml.Unmarshal([]byte(infoStr), &info)
	if err != nil {
		panic(err)
	}
	return &info
}

func (i *Info) GetName() string {
	return i.name
}

func (i *Info) GetDescription() string {
	return i.description
}

func (i *Info) GetPath() string {
	return i.path
}

func (i *Info) GetMethod() interfaces.Method {
	return i.method
}

func (i *Info) GetRequestBodySchema() interfaces.RequestBodySchema {
	return i.requestBodySchema
}

func (i *Info) GetResponseBodySchema() interfaces.ResponseBodySchema {
	return i.responseBodySchema
}

func (i *Info) SetName(name string) {
	i.name = name
}

func (i *Info) SetDescription(description string) {
	i.description = description
}

func (i *Info) SetPath(path string) {
	i.path = path
}

func (i *Info) SetMethod(method interfaces.Method) {
	i.method = method
}

func (i *Info) SetRequestBodySchema(requestBodySchema interfaces.RequestBodySchema) {
	i.requestBodySchema = requestBodySchema
}

func (i *Info) SetResponseBodySchema(responseBodySchema interfaces.ResponseBodySchema) {
	i.responseBodySchema = responseBodySchema
}
