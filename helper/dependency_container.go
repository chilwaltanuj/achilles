package helper

import (
	"achilles/model"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func initializeDependecies(configuration *model.ApplicationConfiguration) {
	globalContainer = dig.New()

	globalContainer.Provide(func() *model.ApplicationConfiguration {
		return configuration
	})
	globalContainer.Provide(func() *logrus.Logger {
		return globalLogger
	})
}

// GetGlobalDependencyContainer returns the globally container.
func GetGlobalDependencyContainer() *dig.Container {
	return globalContainer
}
