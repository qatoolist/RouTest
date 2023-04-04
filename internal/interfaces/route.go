package interfaces

import "net/http"

// Route represents a custom HTTP request with additional fields.
type Route interface {
	Send() (*http.Response, error)
	SetReqBodySchema(schema string) error
	SetResBodySchema(schema string) error
	GetName() string
	GetParentApplication() Application
	GetRouteParametersRegistry() ParametersRegistry
	GetRouteHooksRegistry() HooksRegistry
	GetScenarioRegistry() ScenarioRegistry
	ValidateReqBody(body interface{}) error
	ValidateResBody(body interface{}) error
}
