package models

import "fmt"

// Config is a map of string keys to empty interfaces that represents application configuration.
type Config map[string]interface{}

// Get returns the value of the config key represented by the provided keys.
// Returns an error if the key is not found.
func (c Config) Get(keys ...string) (interface{}, error) {
	var value interface{} = c
	for _, key := range keys {
		v, ok := value.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("config key '%s' not found", key)
		}
		value, ok = v[key]
		if !ok {
			return nil, fmt.Errorf("config key '%s' not found", key)
		}
	}
	return value, nil
}

// Set sets the value of the config key represented by the provided keys to the provided value.
// Creates any necessary sub-maps as needed.
func (c Config) Set(keys []string, value interface{}) {
	m := c
	for _, key := range keys[:len(keys)-1] {
		if _, ok := m[key]; !ok {
			m[key] = make(map[string]interface{})
		}
		m = m[key].(map[string]interface{})
	}
	m[keys[len(keys)-1]] = value
}

func (c Config) GetHost() (*Host, error) {
	hostMap, err := c.Get("host")
	if err != nil {
		return nil, err
	}

	hostValues, ok := hostMap.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid host configuration")
	}

	protocol, ok := hostValues["protocol"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid host configuration: protocol must be a string")
	}

	hostname, ok := hostValues["hostname"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid host configuration: hostname must be a string")
	}

	port, ok := hostValues["port"].(int)
	if !ok {
		return nil, fmt.Errorf("invalid host configuration: port must be a number")
	}

	return NewHost(protocol, hostname, port), nil
}
