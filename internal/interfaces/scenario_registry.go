package interfaces

import "net/http"

type ScenarioRegistry interface {
	// AddScenario adds a scenario to the registry.
	AddScenario(s Scenario)

	// ExportToRequest exports the parameters to the HTTP request.
	ExportToRequest(req *http.Request, scenario Scenario) (*http.Request, error)

	// RunBeforeHooks executes the Before Hooks of all the scenarios in the registry in the order defined in the comment of Scenario struct.
	RunBeforeHooks(scenario Scenario) (Route, error)

	// RunAfterHooks executes the After Hooks of all the scenarios in the registry in the order defined in the comment of Scenario struct.
	RunAfterHooks(scenario Scenario) (Response, error)

	// GetScenarios returns a pointer to the scenarios slice registered in the registry.
	GetScenarios() *[]Scenario
}
