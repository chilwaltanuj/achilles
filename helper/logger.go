package helper

import (
	"achilles/model"
	"fmt"

	"go.uber.org/dig"
)

func InvokeAndLog(container *dig.Container, fn interface{}) {
	if err := container.Invoke(fn); err != nil {
		//log.Error("Error invoking function:", err)
		fmt.Println("Error invoking function:", err)
	}
}

func LogInformation(data interface{}) func(*model.DependencyContainer) {
	return func(dependencyContainer *model.DependencyContainer) {
		dependencyContainer.LogWriter.Info(data)
	}
}

func LogError(data interface{}) func(*model.DependencyContainer) {
	return func(dependencyContainer *model.DependencyContainer) {
		dependencyContainer.LogWriter.Error(data)
	}
}

func LogWarning(data interface{}) func(*model.DependencyContainer) {
	return func(dependencyContainer *model.DependencyContainer) {
		dependencyContainer.LogWriter.Warning(data)
	}
}

func LogFatal(data interface{}) func(*model.DependencyContainer) {
	return func(dependencyContainer *model.DependencyContainer) {
		dependencyContainer.LogWriter.Fatal(data)
	}
}
