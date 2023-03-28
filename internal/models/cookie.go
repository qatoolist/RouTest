package models

import (
	"net/http"
	"time"
)

type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string
	MaxAge     int
	Secure     bool
	HttpOnly   bool
	SameSite   http.SameSite
	Raw        string
	Unparsed   []string
}

type Cookies []*Cookie
