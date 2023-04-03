package interfaces

type Application interface {
	GetRouteByName(name string) (Route, bool)
	AddRoute(name string, route *Route) Route
	LoadRequirements(configPath string) error
	RegisterResponse(name string, resp *Response)
	GetResponse(name string) (*Response, error)
	RegisterParameter(name string, param *Parameter)
	GetParameter(name string) (*Parameter, error)
	GetApplicationParametersRegistry() ParametersRegistry
	GetApplicationHooksRegistry() HooksRegistry
	GetMeta() Meta
	NewRoute(info Info, meta string) Route
}
