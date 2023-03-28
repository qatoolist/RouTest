// Package loaders provides implementations of the ConfigLoader interface
// for loading configuration data from various sources.

package loaders

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/qatoolist/RouTest/models"
)

// YAMLConfigLoader is a ConfigLoader implementation that loads configuration
// data from a YAML file.
type YAMLConfigLoader struct{}

// LoadConfig loads the configuration data for the specified environment from
// a YAML file located at the specified path. The file should be named
// "<env>.yaml". If the file does not exist, an error is returned.
func (l *YAMLConfigLoader) LoadConfig(env string, path string) (*models.Config, error) {
	filename := env + ".yaml"
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

	// Unmarshal the YAML data into the Config map
	var config models.Config
	err = yaml.Unmarshal(contents, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
