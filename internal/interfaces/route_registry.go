package interfaces

type RouteRegistry interface {
	Lock()
	Unlock()
	GetRouteByName(name string) (Route, bool)
	AddRoute(name string, route Route) Route
	Length() int
	GetRegistry() map[string]Route
}
