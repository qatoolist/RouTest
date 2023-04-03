package models

import (
	"fmt"
	"sync"

	"github.com/qatoolist/RouTest/internal/interfaces"
	"github.com/qatoolist/RouTest/internal/loaders"
)

type ConfigImpl struct {
	sync.RWMutex
	config map[string]interface{}
}

func NewConfig() interfaces.Config {
	return &ConfigImpl{
		config: make(map[string]interface{}),
	}
}

func (c *ConfigImpl) Get(keys ...string) (interface{}, error) {
	c.RLock()
	defer c.RUnlock()

	var value interface{} = c.config
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

func (c *ConfigImpl) Set(keys []string, value interface{}) {
	c.Lock()
	defer c.Unlock()

	m := c.config
	for _, key := range keys[:len(keys)-1] {
		if _, ok := m[key]; !ok {
			m[key] = make(map[string]interface{})
		}
		m = m[key].(map[string]interface{})
	}
	m[keys[len(keys)-1]] = value
}

func (c *ConfigImpl) GetHost() (interfaces.Host, error) {
	c.RLock()
	defer c.RUnlock()

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

func (c *ConfigImpl) CopyFromTemp(cnf *loaders.Config) interfaces.Config {
	c.Lock()
	defer c.Unlock()

	temp := make(loaders.Config)
	for key, value := range *cnf {
		temp[key] = value
	}
	for key, value := range temp {
		c.config[key] = value
	}

	return c
}
