// Package loaders provides implementations of the ConfigLoader interface
// for loading configuration data from various sources.

package loaders

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/qatoolist/RouTest/internal/models"
)

// ENVConfigLoader is a ConfigLoader implementation that loads configuration
// data from a .env file.
type ENVConfigLoader struct{}

// LoadConfig loads the configuration data for the specified environment from
// a .env file located at the specified path. The file should be named
// "<env>.env". If the file does not exist, an error is returned.
func (l *ENVConfigLoader) LoadConfig(env string, path string) (*models.Config, error) {
	filename := env + ".env"
	filepath := path + "/" + filename

	// Check if the file exists
	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("file not found")
		}
		return nil, err
	}

	// Load the environment variable configuration data
	err = godotenv.Load(filepath)
	if err != nil {
		return nil, err
	}

	// Parse the environment variable configuration data into the Config map
	config := make(models.Config)
	for _, envVar := range os.Environ() {
		parts := strings.SplitN(envVar, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid environment variable configuration line: %s", envVar)
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		config.Set(strings.Split(key, "."), value)
	}

	return &config, nil
}
