package models

import "github.com/qatoolist/RouTest/internal/interfaces"

// RouteIterator is an iterator over the routes in a RoutesRegistry.
type RouteIterator struct {
	registry interfaces.RouteRegistry
	index    int
}

// NewRouteIterator creates a new iterator for the given RoutesRegistry.
func NewRouteIterator(registry interfaces.RouteRegistry) *RouteIterator {
	return &RouteIterator{
		registry: registry,
		index:    -1, // start at -1 so first call to Next() moves to index 0
	}
}

// Next returns the next Route in the registry, or nil if there are no more routes.
func (i *RouteIterator) Next() interfaces.Route {
	i.index++
	i.registry.Lock()
	defer i.registry.Unlock()
	if i.index >= i.registry.Length() {
		return nil
	}
	var route interfaces.Route
	for _, r := range i.registry.GetRegistry() {
		if i.index == 0 {
			route = r
		} else {
			i.index--
		}
	}
	return route
}
