package interfaces

// RouteIterator provides an iterator over the routes in a Routes map.
type RouteIterator interface {
	Next() (Route, bool)
}
