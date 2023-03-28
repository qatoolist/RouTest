// Package models provides Route handling and other utilities.
package models

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

// Route represents a custom HTTP request with additional fields.
type Route struct {
	Method           string
	URL              *url.URL
	Proto            string
	ProtoMajor       int
	ProtoMinor       int
	Header           Header
	Body             []byte
	Cookies          Cookies
	ContentLength    int64
	TransferEncoding []string
	Close            bool
	Host             string
	Form             url.Values
	PostForm         url.Values
	MultipartForm    *multipart.Form
	Trailer          http.Header
	RemoteAddr       string
	RequestURI       string
	TLS              *tls.ConnectionState
	Cancel           <-chan struct{}
	Response         <-chan *http.Response
	ctx              context.Context
	Deadline         time.Time
	CancelFunc       context.CancelFunc
}

// NewRoute creates a new instance of the Route struct with default values for its fields.
func NewRoute() *Route {
	return &Route{
		Method:        "GET",
		URL:           nil,
		Proto:         "",
		ProtoMajor:    0,
		ProtoMinor:    0,
		Header:        Header{},
		Body:          []byte{},
		Cookies:       []*Cookie{},
		ContentLength: 0,
		Close:         false,
		Host:          "",
		Form:          url.Values{},
		PostForm:      url.Values{},
		MultipartForm: &multipart.Form{},
		Trailer:       http.Header{},
		RemoteAddr:    "",
		RequestURI:    "",
		TLS:           &tls.ConnectionState{},
		Cancel:        nil,
		Response:      nil,
		ctx:           context.Background(),
		Deadline:      time.Time{},
		CancelFunc: func() {
		},
	}
}

// SetMethod sets the HTTP method for the request.
func (r *Route) SetMethod(method string) *Route {
	r.Method = method
	return r
}

// SetURL sets the URL for the request.
func (r *Route) SetURL(url *url.URL) *Route {
	r.URL = url
	return r
}

func (r *Route) SetHeaders(headers http.Header) *Route {
	for key, values := range headers {
		r.Header.Del(key)
		for _, value := range values {
			r.Header.Add(key, value)
		}
	}
	return r
}

// SetCookies sets the HTTP cookies for the request.
func (r *Route) SetCookies(cookies Cookies) *Route {
	r.Cookies = cookies
	return r
}

// SetBody sets the HTTP body for the request.
func (r *Route) SetBody(body []byte) *Route {
	r.Body = body
	return r
}

// AddQueryParam adds a query parameter to the request URL.
func (r *Route) AddQueryParam(key, value string) *Route {
	q := r.URL.Query()
	q.Add(key, value)
	r.URL.RawQuery = q.Encode()
	return r
}

// AddFormValue adds a form value to the request body.
func (r *Route) AddFormValue(key, value string) *Route {
	if r.Body == nil {
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		r.Body = []byte{}
	}
	body := string(r.Body)
	if len(body) > 0 {
		body += "&"
	}
	body += fmt.Sprintf("%s=%s", key, url.QueryEscape(value))
	r.Body = []byte(body)
	return r
}

// AddMultipartFile adds a multipart file to the request body.
func (r *Route) AddMultipartFile(fieldName, fileName string, file io.Reader) *Route {
	if r.MultipartForm == nil {
		r.MultipartForm = &multipart.Form{}
	}
	writer := multipart.NewWriter(r.MultipartForm)
	defer writer.Close()
	part, err := writer.CreateFormFile(fieldName, fileName)
	if err != nil {
		return r
	}
	io.Copy(part, file)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	return r
}

// AddMultipartFormField adds a multipart form field to the request body.
func (r *Route) AddMultipartFormField(fieldName, value string) *Route {
	if r.MultipartForm == nil {
		r.MultipartForm = &multipart.Form{}
	}
	writer := multipart.NewWriter(r.MultipartForm)
	defer writer.Close()
	writer.WriteField(fieldName, value)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	return r
}

// AddJSONBody adds a JSON body to the request.
func (r *Route) AddJSONBody(body interface{}) *Route {
	if r.Body == nil {
		r.Header.Set("Content-Type", "application/json")
		r.Body = []byte{}
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return r
	}
	r.Body = jsonBody
	return r
}

// AddXMLBody adds an XML body to the request.
func (r *Route) AddXMLBody(body interface{}) *Route {
	if r.Body == nil {
		r.Header.Set("Content-Type", "application/xml")
		r.Body = []byte{}
	}
	xmlBody, err := xml.Marshal(body)
	if err != nil {
		return r
	}
	r.Body = xmlBody
	return r
}

// AddTextBody adds a plain text body to the request.
func (r *Route) AddTextBody(body string) *Route {
	if r.Body == nil {
		r.Header.Set("Content-Type", "text/plain")
		r.Body = []byte{}
	}
	r.Body = []byte(body)
	return r
}

func bytesToReader(body []byte) io.Reader {
	return bytes.NewReader(body)
}

// Send sends the HTTP request and returns the HTTP response.
func (r *Route) Send() (*http.Response, error) {
	client := http.DefaultClient
	req, err := http.NewRequest(r.Method, r.URL.String(), bytesToReader(r.Body))
	if err != nil {
		return nil, err
	}
	req.Header = r.Header.ToHTTPHeader()
	for name, value := range r.Cookies {
		req.AddCookie(&http.Cookie{
			Name:       name,
			Value:      value,
			Path:       "",
			Domain:     "",
			Expires:    time.Time{},
			RawExpires: "",
			MaxAge:     0,
			Secure:     false,
			HttpOnly:   false,
			SameSite:   0,
			Raw:        "",
			Unparsed:   []string{},
		})
	}
	return client.Do(req)
}

// SendWithContext sends the HTTP request with a context and returns the HTTP response.
func (r *Route) SendWithContext(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	req, err := http.NewRequestWithContext(ctx, r.Method, r.URL.String(), bytesToReader(r.Body))
	if err != nil {
		return nil, err
	}
	req.Header = r.Header.ToHTTPHeader()
	for _, cookie := range r.Cookies {
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
	return client.Do(req)
}
