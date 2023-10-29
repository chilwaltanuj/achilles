package helper

import (
	"achilles/model"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

var globalLogger *logrus.Logger
var container *dig.Container

func BuildDependencies(configuration *model.ApplicationConfiguration) {
	initializeLogger(configuration.Log)
	initializeDependecies(configuration)
}
