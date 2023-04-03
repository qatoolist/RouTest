// Package models provides types for defining an application.
package models

import (
	"errors"

	"github.com/qatoolist/RouTest/internal/interfaces"
)

// Application represents an HTTP application.
type Application struct {
	// Requirements contains the application's requirements.
	Requirements interfaces.Requirements

	// Meta provides metadata about the application.
	Meta interfaces.Meta

	// Config specifies the application's configuration.
	Config interfaces.Config

	// Host specifies the host for the application.
	Host interfaces.Host

	// Environment specifies the RunEnvironment for the application under test.
	Environment string

	// RouteRegistry specifies the application's routes.
	RouteRegistry interfaces.RouteRegistry

	// ApplicationParametersRegistry are the application level parameters and
	//These Parameters are available through all scenarios of all routes defined under this application
	//But can be overriden by providing the route or scenario level parameters having same keys.
	ApplicationParametersRegistry interfaces.ParametersRegistry

	// ApplicationHooksRegistry is a registry of Before and After Hooks defined at application level
	// The Before Hooks are triggered for every scenario of every route defined under this application
	// The Order of Hooks Execution is -  BeforeApplicationHooks, BeforeRouteHooks, BeforeScenarioHooks
	ApplicationHooksRegistry interfaces.HooksRegistry

	// Register specifies the register of Fixed Parameters and Responses that you many want to
	// Refer later just by using the human redable names of the parameters.
	Register interfaces.Register
}

// NewApplication creates a new Application object.
func NewApplication(env string, config interfaces.Config, requirements interfaces.Requirements, meta interfaces.Meta, host interfaces.Host) (*Application, error) {
	// NewApplication(requirements_path string, meta string, config interfaces.Config)  {

	return &Application{
		Requirements:                  requirements,
		Meta:                          meta,
		Environment:                   env,
		Config:                        config,
		Host:                          host,
		RouteRegistry:                 NewRouteRegistry(),
		ApplicationParametersRegistry: NewParameterRegistry(),
		ApplicationHooksRegistry:      NewHooksRegistry(),
		Register:                      NewRegister(),
	}, nil
}

func (a *Application) NewRoute(info interfaces.Info, meta string) interfaces.Route {
	new_meta, err := NewMetaFromString(meta)
	if err != nil {
		panic(err)
	}
	route := Route{
		Info:                    Info{},
		ParentApplication:       a,
		Meta:                    new_meta,
		ScenarioRegistry:        NewScenarioRegistry(),
		RouteParametersRegistry: NewParameterRegistry(),
		RouteHooksRegistry:      NewHooksRegistry(),
		Body:                    nil,
		Response:                nil,
	}
	return interfaces.Route(&route)
}

// GetRouteByName retrieves a route by its name.
func (a *Application) GetRouteByName(name string) (interfaces.Route, bool) {
	return a.RouteRegistry.GetRouteByName(name)
}

// AddRoute adds a new route to the application.
func (a *Application) AddRoute(name string, route *interfaces.Route) interfaces.Route {
	return a.RouteRegistry.AddRoute(name, *route)
}

// LoadRequirements loads the requirements from a YAML file located at the specified path.
func (app *Application) LoadRequirements(reqPath string) error {

	requirements, err := NewRequirementsFromPath(reqPath)
	if err != nil {
		return err
	}

	app.Requirements = requirements
	return nil
}

// RegisterResponse registers a new response to the application register with the given name.
func (app *Application) RegisterResponse(name string, resp *interfaces.Response) {
	app.Register.AddResponse(name, resp)
}

// GetResponse retrieves the response with the given name.
func (app *Application) GetResponse(name string) (*interfaces.Response, error) {
	return app.Register.GetResponse(name)
}

// RegisterParameter registers a new parameter to the application register with the given name.
func (app *Application) RegisterParameter(name string, param *interfaces.Parameter) {
	app.Register.AddParameter(name, param)
}

// GetParameter retrieves the parameter with the given name.
func (app *Application) GetParameter(name string) (*interfaces.Parameter, error) {
	return app.Register.GetParameter(name)
}

// GetParameter retrieves the parameter with the given name.
func (app *Application) GetMeta() interfaces.Meta {
	return app.Meta.Copy()
}

// GetParametersRegistry returns the application-level ParametersRegistry.
func (app *Application) GetApplicationParametersRegistry() interfaces.ParametersRegistry {
	return app.ApplicationParametersRegistry
}

// GetApplicationHooksRegistry returns the application-level HooksRegistry.
func (app *Application) GetApplicationHooksRegistry() interfaces.HooksRegistry {
	return app.ApplicationHooksRegistry
}

// GetScenarios retrieves all Scenarios for the application.
func (app *Application) GetScenarios() []interfaces.Scenario {
	var scenarios []interfaces.Scenario
	for _, route := range app.RouteRegistry.GetRegistry() {
		scenarios = append(scenarios, *route.GetScenarioRegistry().GetScenarios()...)
	}
	return scenarios
}

// GetScenariosByRoute retrieves all Scenarios for a Route by name and returns an error if not found.
func (app *Application) GetScenariosByRoute(name string) ([]interfaces.Scenario, error) {
	route, ok := app.GetRouteByName(name)
	if !ok {
		return nil, errors.New("Route not found")
	}
	return *route.GetScenarioRegistry().GetScenarios(), nil
}
