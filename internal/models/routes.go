package models

// Routes represents a mapping of route name to the corresponding Route instance.
type Routes map[string]*Route

// GetRouteByName returns the Route with the given name.
func (r Routes) GetRouteByName(name string) (*Route, bool) {
	route, ok := r[name]
	return route, ok
}

// AddRoute adds the given Route to the Routes map.
func (r Routes) AddRoute(name string, route *Route) *Route {
	r[name] = route
	return route
}
