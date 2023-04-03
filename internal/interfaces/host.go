package interfaces

// Host represents the protocol, hostname and port of a server.
type Host interface {
	// BaseURL returns the base URL of the server.
	BaseURL() string
	Protocol() string
	Hostname() string
	Port() int
}
