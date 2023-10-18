package main

import (
	"achilles/client"
	"achilles/config"
	"achilles/model"
	"achilles/route"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/dig"
)

func main() {

	var container = dig.New()
	if err := container.Provide(BuildAndGetDependencyContainer); err != nil {
		fmt.Printf("Failed to provide ApplicationDependencies: %s\n", err)
		fmt.Printf("waah" + err.Error())
		return
	}
	//BuildAndGetDependencyContainer()
	initiateAndBuildServer(container)
}

func BuildAndGetDependencyContainer() (*model.DependencyContainer, error) {
	dependency := model.DependencyContainer{}
	var err error
	if dependency.ApplicationConfiguration, err = config.BuildAndGetApplicationConfiguration(); err != nil {
		return nil, err
	}
	dependency.LogWriter = client.BuildAndGetLogWriter(dependency.ApplicationConfiguration.Log)

	return &dependency, nil
}

func initiateAndBuildServer(container *dig.Container) {
	router, _ := route.SetupRouter(container)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: router,
	}

	if err := startServer(server); err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
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

	fmt.Printf("Received signal: %v\n", sig)
	fmt.Println("I cannot believe you canceled on me! :( Achilles sad ")

	if err := server.Shutdown(nil); err != nil {
		fmt.Printf("Failed to gracefully shutdown server: %s\n", err)
	}
}
