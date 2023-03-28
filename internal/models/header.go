package models

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Header represents a custom HTTP header.
type Header map[string][]string

// Add adds a new value to the header for a given key.
func (h Header) Add(key, value string) {
	h[key] = append(h[key], value)
}

// Get returns the first value of a given header key.
func (h Header) Get(key string) string {
	if values, ok := h[key]; ok && len(values) > 0 {
		return values[0]
	}
	return ""
}

// Set sets the value of a given header key.
func (h Header) Set(key, value string) {
	h[key] = []string{value}
}

// Del deletes a given header key.
func (h Header) Del(key string) {
	delete(h, key)
}

// Write writes the headers to an HTTP request or response.
func (h Header) Write(w io.Writer) error {
	for key, values := range h {
		for _, value := range values {
			if _, err := fmt.Fprintf(w, "%s: %s\r\n", key, value); err != nil {
				return err
			}
		}
	}
	return nil
}

// Read reads the headers from an HTTP request or response.
func (h Header) Read(r io.Reader) error {
	var key, value string
	var err error
	var done bool

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			done = true
			break
		}
		if strings.Contains(line, ":") {
			parts := strings.SplitN(line, ":", 2)
			key = strings.TrimSpace(parts[0])
			value = strings.TrimSpace(parts[1])
			h.Add(key, value)
		} else if key != "" {
			value = strings.TrimSpace(line)
			h.Add(key, value)
		}
	}

	if !done {
		if err = scanner.Err(); err != nil {
			return err
		} else {
			return errors.New("missing end of headers")
		}
	}

	return nil
}

func (h Header) ToHTTPHeader() http.Header {
	header := http.Header{}
	for key, values := range h {
		for _, value := range values {
			header.Add(key, value)
		}
	}
	return header
}
