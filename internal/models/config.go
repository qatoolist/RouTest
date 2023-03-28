package models

import "fmt"

type Config map[string]interface{}

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

func (c Config) Set(keys []string, value interface{}) {
	m := c
	for i, key := range keys[:len(keys)-1] {
		if _, ok := m[key]; !ok {
			m[key] = make(map[string]interface{})
		}
		m = m[key].(map[string]interface{})
	}
	m[keys[len(keys)-1]] = value
}
