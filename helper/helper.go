package helper

import (
	"achilles/client"
	"achilles/constant"
	"achilles/model"
	"context"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

//TODO : Revisit global variable and DI container post intergration of few clients e.g. http , postgres etc.

var globalLogger *logrus.Logger
var clientHTTP *model.ClientHttp
var globalContainer *dig.Container
var globalConfiguration *model.ApplicationConfiguration

func BuildDependencies(appConfiguration *model.ApplicationConfiguration) {
	globalConfiguration = appConfiguration
	globalLogger = client.BuildAndGetLogger(appConfiguration.Log)
	clientHTTP = client.NewHTTPClient(ApplicationConfiguration().HTTP, GetLogger())
	initializeDependecies(appConfiguration)

	LogDetails(logrus.InfoLevel, constant.DependenciesLoaded, *globalConfiguration)
}

func Execute[T any](ginContext context.Context, request model.ClientHTTPRequest) model.ClientResponseDetails {
	return client.Execute[T](ginContext, clientHTTP, request)
}

func ApplicationConfiguration() *model.ApplicationConfiguration {
	return globalConfiguration
}

func GetHttpClient() *model.ClientHttp {
	return clientHTTP
}

func GetLogger() *logrus.Logger {
	return globalLogger
}
