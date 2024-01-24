package client

import (
	"achilles/model"
	"errors"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// TODO - refactoring needed
func NewHTTPClient(config model.ClientHTTPConfig, logger *logrus.Logger) *model.ClientHttp {
	if config.RequestTimeoutDuration <= 0 || config.RetryCountMax < 0 || config.MaxConcurrentRequests <= 0 {
		logger.Error("Invalid configuration for HTTP client")
		return nil
	}

	restyClient := resty.New()

	restyClient.SetTimeout(config.RequestTimeoutDuration).
		SetRetryCount(config.RetryCountMax).
		SetRetryWaitTime(config.RetryBackoffDuration).
		SetRetryMaxWaitTime(config.RetryMaxWaitDuration).
		SetLogger(logger).
		SetDebug(true). // Assuming DebugMode is a boolean in ClientHTTPConfig
		SetRetryAfter(func(client *resty.Client, resp *resty.Response) (time.Duration, error) {
			if resp.StatusCode() == http.StatusTooManyRequests {
				return time.Duration(config.RetryBackoffDuration), nil
			}
			return 0, errors.New("quota exceeded")
		}).
		OnError(func(req *resty.Request, err error) {
			logger.Errorf("HTTP request error: %v, URL: %s, Method: %s", err, req.URL, req.Method)
		}).
		OnBeforeRequest(func(client *resty.Client, req *resty.Request) error {
			logger.Infof("Sending request: URL: %s, Method: %s", req.URL, req.Method)
			return nil
		}).
		OnAfterResponse(func(client *resty.Client, resp *resty.Response) error {
			logger.Infof("Received response: Status: %s, Time: %v", resp.Status(), resp.Time())
			return nil
		})

	hystrix.ConfigureCommand(config.CircuitBreakerName, hystrix.CommandConfig{
		Timeout:                int(config.RequestTimeoutDuration / time.Millisecond),
		MaxConcurrentRequests:  config.MaxConcurrentRequests,
		ErrorPercentThreshold:  config.ErrorThresholdPercentage,
		RequestVolumeThreshold: config.RequestVolumeThreshold,
		SleepWindow:            config.CircuitBreakerActiveTimeMs,
	})

	return &model.ClientHttp{Client: restyClient, Logger: logger}
}
