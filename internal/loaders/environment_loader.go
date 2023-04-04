package loaders

import (
	"errors"
	"os"
)

// LoadRoutestsEnv loads the Routests environment from the ROUTESTS_ENV environment variable.
func LoadRoutestsEnv() (string, error) {
	env := os.Getenv("ROUTESTS_ENV")
	if env == "" {
		return "", errors.New("ROUTESTS_ENV environment variable not set")
	}
	return env, nil
}
