package models

import (
	"fmt"

	"github.com/qatoolist/RouTest/internal/interfaces"
)

// Host represents the protocol, hostname and port of a server.
type Host struct {
	// The protocol to use, e.g. "http" or "https"
	protocol string

	// The hostname of the server
	hostname string

	// The port number of the server
	port int
}

// NewHost creates a new Host object with the given protocol, hostname and port.
func NewHost(protocol string, hostname string, port int) interfaces.Host {
	return &Host{
		protocol: protocol,
		hostname: hostname,
		port:     port,
	}
}

// BaseURL returns the base URL of the server.
func (h *Host) BaseURL() string {
	protocol := h.Protocol()
	if protocol == "" {
		protocol = "http"
	}
	port := ""
	if h.Port() != 0 {
		port = fmt.Sprintf(":%d", h.Port())
	}
	return fmt.Sprintf("%s://%s%s", protocol, h.Hostname(), port)
}

// BaseURL returns the base URL of the server.
func (h *Host) Protocol() string {
	return h.protocol
}

// BaseURL returns the base URL of the server.
func (h *Host) Hostname() string {
	return h.hostname
}

// BaseURL returns the base URL of the server.
func (h *Host) Port() int {
	return h.port
}
