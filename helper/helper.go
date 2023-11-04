package helper

import (
	"achilles/constant"
	"achilles/model"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

var globalLogger *logrus.Logger
var globalContainer *dig.Container
var globalConfiguration *model.ApplicationConfiguration

func BuildDependencies(appConfiguration *model.ApplicationConfiguration) {
	globalConfiguration = appConfiguration
	globalLogger = buildAndGetLogger(appConfiguration.Log)
	initializeDependecies(appConfiguration)

	LogDetails(logrus.InfoLevel, constant.DependenciesLoaded, *globalConfiguration)
}
