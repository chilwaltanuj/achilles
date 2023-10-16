package model

import "github.com/sirupsen/logrus"

type DependencyContainer struct { // DependencyConatiner
	ApplicationConfiguration *ApplicationConfiguration

	LogWriter *logrus.Logger
}
