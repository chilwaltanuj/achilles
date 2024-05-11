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
var clientRDBMS *model.ClientRDBMS
var globalContainer *dig.Container
var globalConfiguration *model.ApplicationConfiguration
var blockerError error

func BuildDependencies(appConfiguration *model.ApplicationConfiguration) error {
	globalConfiguration = appConfiguration
	globalLogger = client.BuildAndGetLogger(appConfiguration.Log)

	if clientHTTP, blockerError = client.NewHTTPClient(ApplicationConfiguration().HTTP, GetLogger()); blockerError != nil {
		return blockerError
	}
	if clientRDBMS, blockerError = client.NewRDBMSClient(ApplicationConfiguration().RDBMS, GetLogger()); blockerError != nil {
		return blockerError
	}

	//initializeDependecies(appConfiguration)

	LogDetails(constant.LogLevelInfo, constant.DependenciesLoaded, *globalConfiguration)

	return nil
}

func ExecuteHttpRequest[T any](ginContext context.Context, request model.ClientHTTPRequest) model.ClientResponseDetails {
	return client.Execute[T](ginContext, clientHTTP, request)
}

func ExecuteRdbmsQuery[T any](ginContext context.Context, request model.RequestClientRDBMS) ([]T, error) {
	return client.ExecuteQuery[T](ginContext, clientRDBMS, request)
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
