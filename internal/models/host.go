package models

import "fmt"

// Host represents the protocol, hostname and port of a server.
type Host struct {
	// The protocol to use, e.g. "http" or "https"
	Protocol string

	// The hostname of the server
	Hostname string

	// The port number of the server
	Port int
}

// NewHost creates a new Host object with the given protocol, hostname and port.
func NewHost(protocol string, hostname string, port int) *Host {
	return &Host{
		Protocol: protocol,
		Hostname: hostname,
		Port:     port,
	}
}

// BaseURL returns the base URL of the server.
func (h *Host) BaseURL() string {
	protocol := h.Protocol
	if protocol == "" {
		protocol = "http"
	}
	port := ""
	if h.Port != 0 {
		port = fmt.Sprintf(":%d", h.Port)
	}
	return fmt.Sprintf("%s://%s%s", protocol, h.Hostname, port)
}
