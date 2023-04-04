package loaders

type Config map[string]interface{}

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
