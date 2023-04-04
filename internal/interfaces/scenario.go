package interfaces

type Scenario interface {
	// GetParentRoute returns the reference to the parent route.
	GetParentRoute() Route

	// GetInfo returns Information about the scenario.
	GetInfo() Info

	// GetMeta returns metadata about the scenario.
	GetMeta() *Meta

	// GetRequestBodySchema returns the schema for the request body.
	GetRequestBodySchema() *RequestBodySchema

	// GetResponseBodySchema returns the schema for the response body.
	GetResponseBodySchema() *ResponseBodySchema

	// GetScenarioParametersRegistry returns the scenario level parameters and
	// The list of parameters is derived from the route level parameters
	// and the route level parameters are always available through the scope of this scenario
	// These Parameters are available only for the request being sent as part of this scenario
	// But can be overriden by providing the route or scenario level parameters having same keys.
	GetScenarioParametersRegistry() ParametersRegistry

	// GetScenarioHooksRegistry returns the registry of Before and After Hooks defined at scenario level
	// The Before Hooks are triggered for Before the request is being sent in the scope of this scenario
	// The After Hooks are triggered for After the response has been receieved in the scope of this scenario
	// The Order of Before Hooks Execution is -  BeforeApplicationHooks, BeforeRouteHooks, BeforeScenarioHooks
	// The Order of After Hooks Execution is -  AfterApplicationHooks, AfterRouteHooks, AfterScenarioHooks
	GetScenarioHooksRegistry() HooksRegistry

	// GetResponse returns the HTTP Response received after sending the request for this scenario
	GetResponse() Response
}
