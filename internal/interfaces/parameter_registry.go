package interfaces

import "net/http"

type ParametersRegistry interface {
	RegisterQueryParameter(key string, value string) error
	RegisterPathVariable(key string, value string) error
	RegisterHeader(key string, value string) error
	GetQueryParameters() []Parameter
	GetPathVariables() []Parameter
	GetHeaders() []Parameter
	ExportToRequest(req *http.Request) (*http.Request, error)
	ImportFromHTTPResponse(httpResp *http.Response) error
	GetParameterByKey(key string, pType string) (string, error)
}
