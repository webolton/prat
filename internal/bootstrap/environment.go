package bootstrap

import (
	"errors"
	"log/slog"
	"os"
	"slices"
)

func getEnvironment() (string, error) {
	var validEnvironments = []string{"development", "test", "staging", "production"}
	environment := os.Getenv("Environment")

	if environment == "" {
		environment = "development"
	}

	if slices.Contains(validEnvironments, environment) {
		return environment, nil
	}

	return "", errors.New("invalid environment")
}

func setEnvironment() {
	if appEnv, err := getEnvironment(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	} else {
		appConfig.environment = appEnv
	}
}
