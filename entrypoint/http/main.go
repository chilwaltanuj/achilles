package main

import (
	"achilles/config"
	"achilles/constant"
	"achilles/helper"
	"achilles/model"
	"achilles/route"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	configuration, err := config.PrepareAndFetchApplicationConfiguration()
	if err != nil {
		fmt.Print(constant.ServerStartFailure, err)
		return
	}

	helper.BuildDependencies(configuration)
	BuildServer(configuration)
}

func BuildServer(configuration *model.ApplicationConfiguration) {
	server := getServer(configuration)
	go startServer(server)
	gracefulShutdownOnClosureSignals(server)
}

func startServer(server *http.Server) {
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		helper.LogDetails(constant.LogLevelFatal, constant.ServerStartFailure, err)
		os.Exit(1)
	}
}

func getServer(configuration *model.ApplicationConfiguration) *http.Server {
	router, _ := route.SetupRouter()

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", configuration.HttpPort),
		Handler: router,
	}
}

// gracefulShutdownOnClosureSignals listens to user-initiated closure (Ctrl+C) and process kill (pkill processID)
// and performs graceful shutdown when a signal is received.
func gracefulShutdownOnClosureSignals(server *http.Server) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	sig := <-signalChannel

	message := fmt.Sprintf(constant.ServerShutdownRequest, sig)
	helper.LogDetails(constant.LogLevelFatal, message, "")

	if err := server.Shutdown(context.TODO()); err != nil {
		helper.LogDetails(constant.LogLevelWarn, constant.GracefulShutdownError, err)
	} else {
		helper.LogDetails(constant.LogLevelInfo, constant.GracefulShutdownSuccess, nil)
	}
}
