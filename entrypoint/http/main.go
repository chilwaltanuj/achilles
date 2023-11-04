package main

import (
	"achilles/config"
	"achilles/constant"
	"achilles/helper"
	"achilles/model"
	"achilles/route"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	configuration, err := config.BuildAndGetApplicationConfiguration()
	if err != nil {
		fmt.Print(constant.FailedServerStart, err)
		return
	}

	helper.BuildDependencies(configuration)
	BuildServer(configuration)
}

func BuildServer(configuration *model.ApplicationConfiguration) {
	router, _ := route.SetupRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", configuration.HttpPort),
		Handler: router,
	}

	if err := startServer(server); err != nil {
		helper.LogDetails(logrus.FatalLevel, constant.FailedServerStart, err)
	} else {
		gracefulShutdownOnClosureSignals(server)
	}
}

// StartServer starts the HTTP server in a goroutine.
func startServer(server *http.Server) error {
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			helper.LogDetails(logrus.FatalLevel, constant.FailedServerStart, err)
		}
	}()

	return nil
}

// gracefulShutdownOnClosureSignals listens to user-initiated closure (Ctrl+C) and process kill (pkill processID)
// and performs graceful shutdown when a signal is received.
func gracefulShutdownOnClosureSignals(server *http.Server) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	sig := <-signalChannel

	message := fmt.Sprintf(constant.ServerShutdownRequest, sig)
	helper.LogDetails(logrus.FatalLevel, message, "")

	if err := server.Shutdown(nil); err != nil {
		helper.LogDetails(logrus.ErrorLevel, constant.GracefulShutdownError, err)
	}
}
