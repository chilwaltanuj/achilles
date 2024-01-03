package client

import (
	"achilles/model"

	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/hystrix"
	"github.com/sirupsen/logrus"
)

// NewHTTPClient creates a new hystrix client with the given configuration.
func NewHTTPClient(config model.ClientHTTPConfig, logger *logrus.Logger) model.ClientHttp {
	hystrixClient := hystrix.NewClient(
		hystrix.WithCommandName(config.HystrixCommand),
		hystrix.WithHTTPTimeout(config.TimeoutDuration),
		hystrix.WithHystrixTimeout(config.TimeoutDuration),
		hystrix.WithMaxConcurrentRequests(config.MaxConcurrent),
		hystrix.WithRetrier(heimdall.NewRetrier(heimdall.NewConstantBackoff(config.RetryBackoffDuration, config.RetryJitterDuration))),
		hystrix.WithRetryCount(config.RetryMax),
		hystrix.WithErrorPercentThreshold(config.ErrorThreshold),
		hystrix.WithSleepWindow(config.CircuitBreakerActiveTimeInMs),
	)
	return model.ClientHttp{Client: hystrixClient, Logger: logger}
}
