package client

import (
	"achilles/model"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

// Execute performs a Hystrix-wrapped HTTP request and returns the response.
func Execute[T any](ctx context.Context, clientHttp *model.ClientHttp, req model.ClientHTTPRequest) model.ClientResponseDetails {
	startTime := time.Now()

	// Perform the Hystrix-wrapped HTTP request
	responseObject, hystrixErr := doHystrixRequest[T](ctx, clientHttp, req)

	// Prepare and return the response details
	return prepareResponse[T](hystrixErr, responseObject, startTime)
}

// doHystrixRequest executes the Hystrix-wrapped HTTP request.
func doHystrixRequest[T any](ctx context.Context, clientHttp *model.ClientHttp, req model.ClientHTTPRequest) (*T, error) {
	responseObject := new(T)
	output := make(chan *T, 1)
	errors := hystrix.GoC(ctx, req.HystrixCommand, func(ctx context.Context) error {
		res, err := clientHttp.Client.R().
			SetContext(ctx).
			SetHeaders(req.Headers).
			SetBody(req.Body).
			Execute(req.Method, req.URL)
		if err != nil {
			return err
		}
		err = json.Unmarshal(res.Body(), responseObject)
		if err != nil {
			return err
		}
		output <- responseObject
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// prepareResponse prepares and returns the client response details.
func prepareResponse[T any](hystrixErr error, responseObject *T, startTime time.Time) model.ClientResponseDetails {
	elapsedTime := float32(time.Since(startTime).Microseconds()) / 1000

	status := http.StatusOK
	if hystrixErr != nil {
		// This could be refined to extract actual status code from the error or response object
		status = http.StatusFailedDependency
	}

	return model.ClientResponseDetails{
		IsSuccessful:   hystrixErr == nil,
		Status:         status,
		ResponseStruct: responseObject,
		LatencyInMs:    elapsedTime,
		Error:          hystrixErr,
	}
}
