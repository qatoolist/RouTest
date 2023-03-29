// Package models provides types for defining an application.
package models

// Application represents an HTTP application.
type Application struct {
	// Requirements contains the application's requirements.
	Requirements Requirements

	// Meta provides metadata about the application.
	Meta *Meta

	// Routes specifies the application's routes.
	Routes Routes

	// Config specifies the application's configuration.
	Config Config

	// Host specifies the host for the application.
	Host Host

	// Environment specified the RunEnvironment for the application under test
	Environment string
}

func NewApplication(meta *Meta) *Application {
	return &Application{
		Requirements: Requirements{},
		Meta:         meta.Copy(),
		Routes:       make(Routes, 0),
		Config:       Config{},
		Host:         Host{},
		Environment:  "",
	}
}

// GetRouteByPath retrieves a route by its path.
func (a *Application) GetRouteByName(name string) (*Route, bool) {
	return a.Routes.GetRouteByName(name)
}

// GetRouteByPath retrieves a route by its path.
func (a *Application) AddRoute(name string, route *Route) *Route {
	return a.Routes.AddRoute(name, route)
}
