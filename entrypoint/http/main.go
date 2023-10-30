package main

import (
	"achilles/config"
	"achilles/helper"
	"achilles/model"
	"achilles/route"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	configuration, err := config.BuildAndGetApplicationConfiguration()
	if err != nil {
		fmt.Printf("Failed to load configuration. Reason being -> %v ", err)
		return
	}

	helper.BuildDependencies(configuration)
	initiateAndBuildServer(configuration)
}

func initiateAndBuildServer(configuration *model.ApplicationConfiguration) {
	router, _ := route.SetupRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", configuration.HttpPort),
		Handler: router,
	}

	if err := startServer(server); err != nil {
		helper.GetGlobalLogger().Fatal("Failed to start server. Reason Being" + err.Error())
	} else {
		gracefulShutdownOnClosureSignals(server)
	}
}

// StartServer starts the HTTP server in a goroutine.
func startServer(server *http.Server) error {
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Failed to start server: %s\n", err)
		}
	}()

	return nil
}

// gracefulShutdownOnClosureSignals listens to user-initiated closure (Ctrl+C) and process kill (pkill processID)
// and performs graceful shutdown when a signal is received.
func gracefulShutdownOnClosureSignals(server *http.Server) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	// Wait for a signal to be received
	sig := <-signalChannel

	fmt.Printf("\n Received signal: %v.  I cannot believe you canceled on me! :( Achilles sad ", sig)

	if err := server.Shutdown(nil); err != nil {
		fmt.Printf("Failed to gracefully shutdown server: %s", err)
	}
}
