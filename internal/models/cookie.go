package models

import (
	"net/http"
	"time"
)

// Cookie represents an HTTP cookie.
type Cookie struct {
	// Name specifies the name of the cookie.
	Name string

	// Value specifies the value of the cookie.
	Value string

	// Path specifies the URL path prefix for the cookie.
	Path string

	// Domain specifies the domain name for the cookie.
	Domain string

	// Expires specifies the time when the cookie should expire.
	Expires time.Time

	// RawExpires specifies the cookie expiration time as a string.
	RawExpires string

	// MaxAge specifies the maximum age of the cookie in seconds.
	MaxAge int

	// Secure indicates whether the cookie should only be sent over HTTPS.
	Secure bool

	// HttpOnly indicates whether the cookie should be accessible only via HTTP(S), and not JavaScript.
	HttpOnly bool

	// SameSite sets the SameSite attribute for the cookie. This can be "Strict", "Lax", or "None".
	SameSite http.SameSite

	// Raw specifies the unparsed cookie header string.
	Raw string

	// Unparsed specifies the list of unparsed attribute-value pairs in the cookie header string.
	Unparsed []string
}

type Cookies []*Cookie

// AddCookiesToRequest adds cookies from a Cookies object to an HTTP request.
func (c Cookies) AddCookiesToRequest(req *http.Request) *http.Request {
	for _, cookie := range c {
		req.AddCookie(&http.Cookie{
			Name:       cookie.Name,
			Value:      cookie.Value,
			Path:       cookie.Path,
			Domain:     cookie.Domain,
			Expires:    cookie.Expires,
			RawExpires: cookie.RawExpires,
			MaxAge:     cookie.MaxAge,
			Secure:     cookie.Secure,
			HttpOnly:   cookie.HttpOnly,
			SameSite:   cookie.SameSite,
			Raw:        cookie.Raw,
			Unparsed:   cookie.Unparsed,
		})
	}
	return req
}
