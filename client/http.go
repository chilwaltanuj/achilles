package client

import (
	"achilles/model"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// NewHTTPClient initializes a new HTTP client with given configuration and logger.
func NewHTTPClient(config model.ClientHTTPConfig, logger *logrus.Logger) (*model.ClientHttp, error) {
	if err := validateClientConfig(config); err != nil {
		return nil, fmt.Errorf("HTTP Validation Failed")
	}

	restyClient := resty.New().
		SetTimeout(config.RequestTimeoutDuration).
		SetRetryCount(config.RetryCountMax).
		SetRetryWaitTime(config.RetryBackoffDuration).
		SetRetryMaxWaitTime(config.RetryMaxWaitDuration).
		SetLogger(logger).
		SetDebug(true).
		SetRetryAfter(retryAfterFunc(config)).
		OnError(logRequestError(logger)).
		OnBeforeRequest(logOnBeforeRequestStart(logger)).
		OnAfterResponse(logOnAfterResponseReceived(logger))

	configureHystrix(config)

	return &model.ClientHttp{Client: restyClient, Logger: logger}, nil
}

// validateClientConfig checks if the provided HTTP client configuration is valid.
func validateClientConfig(config model.ClientHTTPConfig) error {
	if config.RequestTimeoutDuration <= 0 {
		return errors.New("invalid request timeout duration")
	}
	if config.RetryCountMax < 0 {
		return errors.New("invalid retry count max")
	}
	if config.MaxConcurrentRequests <= 0 {
		return errors.New("invalid max concurrent requests")
	}
	return nil
}

// logRequestStart logs the start of an HTTP request.
func logOnBeforeRequestStart(logger *logrus.Logger) resty.RequestMiddleware {
	return func(client *resty.Client, req *resty.Request) error {
		logger.Infof("Starting request: URL: %s, Method: %s", req.URL, req.Method)
		return nil
	}
}

// logRequestError logs the details of any errors that occur during a request.
func logRequestError(logger *logrus.Logger) resty.ErrorHook {
	return func(req *resty.Request, err error) {
		logger.Errorf("HTTP request error: %v, URL: %s, Method: %s", err, req.URL, req.Method)
	}
}

// logResponseReceived logs details of the received response.
func logOnAfterResponseReceived(logger *logrus.Logger) resty.ResponseMiddleware {
	return func(client *resty.Client, resp *resty.Response) error {
		logger.Infof("Received response: Status: %s, Time: %v", resp.Status(), resp.Time())
		return nil
	}
}

// retryAfterFunc creates a function to calculate retry duration based on the response.
func retryAfterFunc(config model.ClientHTTPConfig) resty.RetryAfterFunc {
	return func(client *resty.Client, resp *resty.Response) (time.Duration, error) {
		if resp.StatusCode() == http.StatusTooManyRequests {
			return time.Duration(config.RetryBackoffDuration), nil
		}
		return 0, errors.New("request failed")
	}
}

// configureHystrix sets up the hystrix command with the given configuration.
func configureHystrix(config model.ClientHTTPConfig) {
	hystrix.ConfigureCommand(config.CircuitBreakerName, hystrix.CommandConfig{
		Timeout:                int(config.RequestTimeoutDuration / time.Millisecond),
		MaxConcurrentRequests:  config.MaxConcurrentRequests,
		ErrorPercentThreshold:  config.ErrorThresholdPercentage,
		RequestVolumeThreshold: config.RequestVolumeThreshold,
		SleepWindow:            config.CircuitBreakerActiveTimeMs,
	})
}
