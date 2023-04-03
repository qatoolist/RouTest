package models

import (
	"net/http"
	"sync"

	"github.com/qatoolist/RouTest/internal/interfaces"
)

// Scenario defines an HTTP scenario.
type Scenario struct {
	// ParentRoute provides the reference to the parent route.
	ParentRoute interfaces.Route

	// Info provides Information about the scenario.
	Info interfaces.Info

	// Meta provides metadata about the scenario.
	Meta interfaces.Meta

	// RequestBodySchema specifies the schema for the request body.
	RequestBodySchema interfaces.RequestBodySchema

	// ResponseBodySchema specifies the schema for the response body.
	ResponseBodySchema interfaces.ResponseBodySchema

	// ScenarioParametersRegistry are the scenario level parameters and
	// The list of parameters is derived from the route level parameters
	// and the route level parameters are always available through the scope of this scenario
	// These Parameters are available only for the request being sent as part of this scenario
	// But can be overriden by providing the route or scenario level parameters having same keys.
	ScenarioParametersRegistry interfaces.ParametersRegistry

	// ScenarioHooksRegistry is a registry of Before and After Hooks defined at scenario level
	// The Before Hooks are triggered for Before the request is being sent in the scope of this scenario
	// The After Hooks are triggered for After the response has been receieved in the scope of this scenario
	// The Order of Before Hooks Execution is -  BeforeApplicationHooks, BeforeRouteHooks, BeforeScenarioHooks
	// The Order of After Hooks Execution is -  AfterApplicationHooks, AfterRouteHooks, AfterScenarioHooks
	ScenarioHooksRegistry interfaces.HooksRegistry

	// Response represents the HTTP Response received after sending the request for this scenario
	Response interfaces.Response
}

// AddScenario adds a scenario to the registry.
func (s *Scenario) GetParentRoute() interfaces.Route {
	return s.ParentRoute
}

// ScenarioRegistryImpl represents the implementation of the ScenarioRegistry interface.
type ScenarioRegistryImpl struct {
	mux       sync.Mutex
	scenarios []interfaces.Scenario
}

// NewScenarioRegistry creates a new ScenarioRegistryImpl and returns it as a ScenarioRegistry interface.
func NewScenarioRegistry() interfaces.ScenarioRegistry {
	return &ScenarioRegistryImpl{
		mux:       sync.Mutex{},
		scenarios: []interfaces.Scenario{},
	}
}

// AddScenario adds a scenario to the registry.
func (s *ScenarioRegistryImpl) AddScenario(scenario interfaces.Scenario) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.scenarios = append(s.scenarios, scenario)
}

// ExportToRequest exports the parameters to the HTTP request.
func (s *ScenarioRegistryImpl) ExportToRequest(req *http.Request, scenario interfaces.Scenario) (*http.Request, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	route := scenario.GetParentRoute()
	application := route.GetParentApplication()

	// Export application parameters
	if application != nil {
		// Dereference the *interfaces.Application pointer to get the actual interface value

		req, err := application.GetApplicationParametersRegistry().ExportToRequest(req)
		if err != nil {
			return req, err
		}
	}

	// Export route parameters
	if route != nil {
		req, err := route.GetRouteParametersRegistry().ExportToRequest(req)
		if err != nil {
			return req, err
		}
	}

	// Export scenario parameters
	req, err := scenario.GetScenarioParametersRegistry().ExportToRequest(req)
	if err != nil {
		return req, err
	}

	return req, nil
}

// RunBeforeHooks executes all the Before hooks defined at the application, route, and scenario level.
// The order of execution is as follows:
// 1. Before application hooks
// 2. Before route hooks
// 3. Before scenario hooks
func (sr *ScenarioRegistryImpl) RunBeforeHooks(scenario interfaces.Scenario) (interfaces.Route, error) {

	route := scenario.GetParentRoute()
	app := route.GetParentApplication()
	appHooksRegistry := app.GetApplicationHooksRegistry()
	routeHooksRegistry := scenario.GetParentRoute().GetRouteHooksRegistry()
	scenarioHooksRegistry := scenario.GetScenarioHooksRegistry()

	// Run Before Application Hooks
	var err error
	route, err = appHooksRegistry.RunBeforeHooks(route)
	if err != nil {
		return nil, err
	}

	// Run Before Route Hooks
	route, err = routeHooksRegistry.RunBeforeHooks(route)
	if err != nil {
		return nil, err
	}

	// Run Before Scenario Hooks
	route, err = scenarioHooksRegistry.RunBeforeHooks(route)
	if err != nil {
		return nil, err
	}

	return route, nil
}

// RunAfterHooks executes all the After hooks defined at the application, route, and scenario level.
// The order of execution is as follows:
// 1. After application hooks
// 2. After route hooks
// 3. After scenario hooks

func (sr *ScenarioRegistryImpl) RunAfterHooks(scenario interfaces.Scenario) (interfaces.Response, error) {

	route := scenario.GetParentRoute()
	app := route.GetParentApplication()
	appHooksRegistry := app.GetApplicationHooksRegistry()
	routeHooksRegistry := scenario.GetParentRoute().GetRouteHooksRegistry()
	scenarioHooksRegistry := scenario.GetScenarioHooksRegistry()

	response := scenario.GetResponse()
	response, err := scenarioHooksRegistry.RunAfterHooks(response)
	if err != nil {
		return response, err
	}

	// Run After Route Hooks
	response, err = routeHooksRegistry.RunAfterHooks(response)
	if err != nil {
		return response, err
	}

	// Run After Application Hooks
	response, err = appHooksRegistry.RunAfterHooks(response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// GetScenarios returns all the scenarios registered in the registry.
func (s *ScenarioRegistryImpl) GetScenarios() *[]interfaces.Scenario {
	s.mux.Lock()
	defer s.mux.Unlock()
	return &s.scenarios
}

// GetScenariosByRoute returns all the scenarios with the matching parent route.
func (s *ScenarioRegistryImpl) GetScenariosByRoute(route interfaces.Route) []interfaces.Scenario {
	s.mux.Lock()
	defer s.mux.Unlock()
	var scenarios []interfaces.Scenario
	for _, scenario := range s.scenarios {
		if scenario.GetParentRoute() == route {
			scenarios = append(scenarios, scenario)
		}
	}
	return scenarios
}
