// Package loaders provides implementations of the ConfigLoader interface
// for loading configuration data from various sources.

package loaders

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/qatoolist/RouTest/internal/models"
)

// JSONConfigLoader is a ConfigLoader implementation that loads configuration
// data from a JSON file.
type JSONConfigLoader struct{}

// LoadConfig loads the configuration data for the specified environment from
// a JSON file located at the specified path. The file should be named
// "<env>.json". If the file does not exist, an error is returned.
func (l *JSONConfigLoader) LoadConfig(env string, path string) (*models.Config, error) {
	filename := env + ".json"
	filepath := path + "/" + filename

	// Check if the file exists
	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("file not found")
		}
		return nil, err
	}

	// Read the file contents
	contents, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into the Config map
	var config models.Config
	err = json.Unmarshal(contents, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
