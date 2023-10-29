package helper

import (
	"achilles/model"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func initializeDependecies(configuration *model.ApplicationConfiguration) {
	container = dig.New()

	container.Provide(func() *model.ApplicationConfiguration {
		return configuration
	})
	container.Provide(func() *logrus.Logger {
		return globalLogger
	})
}

// GetGlobalDependencyContainer returns the globally container.
func GetGlobalDependencyContainer() *dig.Container {
	return container
}
