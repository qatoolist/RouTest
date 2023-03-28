// Package models defines constants for HTTP methods.
package models

// Method represents an HTTP method.
type Method struct {
    // Name is the name of the HTTP method.
    Name string
}

// Common HTTP methods.
var (
    // GET is the HTTP GET method.
    GET = Method{"GET"}

    // POST is the HTTP POST method.
    POST = Method{"POST"}

    // PUT is the HTTP PUT method.
    PUT = Method{"PUT"}

    // PATCH is the HTTP PATCH method.
    PATCH = Method{"PATCH"}

    // DELETE is the HTTP DELETE method.
    DELETE = Method{"DELETE"}

    // OPTIONS is the HTTP OPTIONS method.
    OPTIONS = Method{"OPTIONS"}

    // HEAD is the HTTP HEAD method.
    HEAD = Method{"HEAD"}
)
