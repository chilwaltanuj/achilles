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

// NewRDBMSClient initializes a new RDBMS client with the given configuration and logger.
// Unlike HTTP requests which are stateless and transient, allowing for retries to be managed per request,
// RDBMS connections are persistent. Retries here are primarily focused on establishing a reliable initial
// connection pool. Query retries, in contrast to HTTP retries, must consider transactional safety to prevent
// issues such as data duplication.
func NewRDBMSClient(config model.ClientRDBMSConfig, logger *logrus.Logger) (*model.ClientRDBMS, error) {
	// Validate RDBMS client configuration.
	if err := validateRDBMSClientConfig(config); err != nil {
		return nil, fmt.Errorf("RDBMS validation failed: %w", err)
	}

	var db *sqlx.DB
	var err error

	// Attempt to connect to the database with retries.
	for attempt := 1; attempt <= config.RetryCountMax+1; attempt++ {
		db, err = sqlx.Connect("postgres", config.DSN)
		if err == nil {
			logger.Infof("Successfully connected to RDBMS on attempt %d", attempt)
			break
		}
		logger.Errorf("Attempt %d to connect to RDBMS failed: %v", attempt, err)
		// Sleep before retrying to connect again, unless it's the last attempt.
		if attempt < config.RetryCountMax+1 {
			time.Sleep(config.RetryBackDuration)
		}
	}

	// If unable to connect after retries, return with error.
	if err != nil {
		return nil, fmt.Errorf("after %d attempts, RDBMS connection failed: %w", config.RetryCountMax+1, err)
	}

	// Set database connection pool parameters.
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)

	// Configure circuit breaker settings for the RDBMS client to enhance resilience.
	configureHystrixForRDBMS(config, logger)

	// Return the configured RDBMS client.
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
