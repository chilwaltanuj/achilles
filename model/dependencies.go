package model

import "github.com/sirupsen/logrus"

type ApplicationDependencies struct {
	ApplicationConfiguration *ApplicationConfiguration

	LogWriter *logrus.Logger
}
