package loaders

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// ConfigLoader is an interface that defines the LoadConfig method for loading
// configuration data from various sources.
type ConfigLoader interface {
	LoadConfig(env string, path string) (*Config, error)
}

// ConfigLoaderImpl is an implementation of the ConfigLoader interface that loads
// configuration data from a file with a specific extension.
type ConfigLoaderImpl struct {
	extension string
	loader    ConfigLoader
}

// LoadConfig loads the configuration data for the specified environment from
// a file located at the specified path. The file should be named
// "<env>.json", "<env>.yaml", or "<env>.env" depending on the file format.
// If the file does not exist, an error is returned.
func (l *ConfigLoaderImpl) LoadConfig(env string, path string) (*Config, error) {
	filename := env + l.extension
	filepath := path + "/" + filename

	// Check if the file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return nil, err
	}

	return l.loader.LoadConfig(env, path)
}

// ConfigLoaderFactory is a factory function that creates a ConfigLoader based
// on the available file extensions in the specified directory.
func ConfigLoaderFactory(env string, path string) (ConfigLoader, error) {
	// Determine the file extensions available in the config directory
	var extensions []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			extensions = append(extensions, filepath.Ext(path))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Create a ConfigLoader for each available file extension
	var loaders []ConfigLoader
	for _, ext := range extensions {
		switch ext {
		case ".json":
			loaders = append(loaders, &JSONConfigLoader{})
		case ".yaml", ".yml":
			loaders = append(loaders, &YAMLConfigLoader{})
		case ".env":
			loaders = append(loaders, &ENVConfigLoader{})
		default:
			// Ignore unsupported file extensions
		}
	}

	// Return an error if no supported file extensions were found
	if len(loaders) == 0 {
		return nil, errors.New("no supported file extensions found")
	}

	// Select the appropriate loader based on the RunEnv environment variable
	runEnv := os.Getenv("RunEnv")
	if runEnv == "" {
		return nil, errors.New("RunEnv environment variable not set")
	}
	for _, loader := range loaders {
		switch loader := loader.(type) {
		case *JSONConfigLoader:
			if _, err := os.Stat(path + "/" + runEnv + ".json"); err == nil {
				return &ConfigLoaderImpl{".json", loader}, nil
			}
		case *YAMLConfigLoader:
			if _, err := os.Stat(path + "/" + runEnv + ".yaml"); err == nil {
				return &ConfigLoaderImpl{".yaml", loader}, nil
			}
			if _, err := os.Stat(path + "/" + runEnv + ".yml"); err == nil {
				return &ConfigLoaderImpl{".yml", loader}, nil
			}
		case *ENVConfigLoader:
			if _, err := os.Stat(path + "/" + runEnv + ".env"); err == nil {
				return &ConfigLoaderImpl{".env", loader}, nil
			}
		}
	}

	// If no loader was selected, return an error
	return nil, fmt.Errorf("no loader found for environment %s", runEnv)
}

/* Example -
func main() {
	// Specify the environment and config directory
	env := "dev"
	configDir := "./config"

	// Create a ConfigLoader for the specified environment and directory
	loader, err := loaders.ConfigLoaderFactory(env, configDir)
	if err != nil {
		panic(err)
	}

	// Load the configuration data for the specified environment
	config, err := loader.LoadConfig(env, configDir)
	if err != nil {
		panic(err)
	}

	// Print the configuration data
	fmt.Printf("Configuration for environment %s:\n", env)
	for key, value := range *config {
		fmt.Printf("%s: %v\n", key, value)
	}
}
*/
