package client

import (
	"achilles/model"
	"fmt"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import the database/sql driver of your choice
	"github.com/sirupsen/logrus"
)

// NewRDBMSClient initializes a new RDBMS client with given configuration and logger.
func NewRDBMSClient(config model.ClientRDBMSConfig, logger *logrus.Logger) (*model.ClientRDBMS, error) {
	if err := validateRDBMSClientConfig(config); err != nil {
		return nil, fmt.Errorf("RDBMS Validation Failed")
	}

	db, err := sqlx.Connect("postgres", config.DSN) // Use the appropriate SQL driver
	if err != nil {
		return nil, fmt.Errorf("RDBMS connection Failed: %w", err)
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)

	configureHystrixForRDBMS(config, logger)

	return &model.ClientRDBMS{DB: db, Config: config, Logger: logger}, nil
}

func validateRDBMSClientConfig(config model.ClientRDBMSConfig) error {
	// Similar validation logic as in your HTTP client setup
	// Return an error if any configuration is invalid
	return nil
}

func configureHystrixForRDBMS(config model.ClientRDBMSConfig, logger *logrus.Logger) {
	hystrix.ConfigureCommand(config.CircuitBreakerName, hystrix.CommandConfig{
		Timeout:                int(config.RequestTimeoutDuration / time.Millisecond),
		MaxConcurrentRequests:  config.MaxConcurrentRequests,
		ErrorPercentThreshold:  config.ErrorThresholdPercentage,
		RequestVolumeThreshold: config.RequestVolumeThreshold,
		SleepWindow:            config.CircuitBreakerActiveTimeMs,
	})
	logger.Info("Hystrix circuit breaker configured for RDBMS client")
}

// Additional functions for logging query execution, similar to HTTP client logging
func logQueryExecutionStart(logger *logrus.Logger, query string, args []interface{}) {
	logger.Infof("Executing query: %s, with args: %v", query, args)
}

func logQueryEnding(logger *logrus.Logger, query string, err error) {
	if err != nil {
		logger.Errorf("Query execution failed: %s, error: %v", query, err)
	} else {
		logger.Infof("Query execution Succesful: %s, error: %v", query, err)
	}
}
