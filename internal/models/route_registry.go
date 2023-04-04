package models

import (
	"sync"

	"github.com/qatoolist/RouTest/internal/interfaces"
)

// RoutesRegistry represents a mapping of route name to the corresponding Route instance.
type RouteRegistry struct {
	sync.RWMutex
	registry map[string]interfaces.Route
}

func NewRouteRegistry() interfaces.RouteRegistry {
	return &RouteRegistry{
		registry: make(map[string]interfaces.Route),
	}
}

// GetRouteByName returns the Route with the given name.
func (r *RouteRegistry) GetRouteByName(name string) (interfaces.Route, bool) {
	r.Lock()
	defer r.Unlock()
	route, ok := r.registry[name]
	return route, ok
}

// GetRouteByName returns the Route with the given name.
func (r *RouteRegistry) Length() int {
	return len(r.registry)
}

// AddRoute adds the given Route to the Routes map.
func (r *RouteRegistry) AddRoute(name string, route interfaces.Route) interfaces.Route {
	r.Lock()
	defer r.Unlock()
	r.registry[name] = route
	return route
}

// GetRouteByName returns the Route with the given name.
func (r *RouteRegistry) Lock() {
	r.RLock()
}

// GetRouteByName returns the Route with the given name.
func (r *RouteRegistry) Unlock() {
	r.RUnlock()
}

func (r *RouteRegistry) GetRegistry() map[string]interfaces.Route {
	// create a new map to hold the copy
	copy := make(map[string]interfaces.Route)

	// iterate over the original map and copy the values
	for k, v := range r.registry {
		copy[k] = v
	}

	// return the copy
	return copy
}
