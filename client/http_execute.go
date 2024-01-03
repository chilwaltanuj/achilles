package client

import (
	"achilles/model"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// Execute performs an HTTP request using the hystrix client, adjusted as per loaded configuration.
func Execute[T any](ginContext context.Context, clientHTTP model.ClientHttp, request model.ClientHTTPRequest) model.ClientResponseDetails {
	isSuccessful := false
	status := http.StatusFailedDependency
	startTime := time.Now() // Start time

	response, err := execute[T](ginContext, clientHTTP, request)
	elapsedTime := float32(time.Since(startTime).Microseconds()) / 1000

	if err == nil {
		isSuccessful = true
		status = http.StatusOK
	}
	clientResponse := model.ClientResponseDetails{
		IsSuccessful:   isSuccessful,
		Status:         status,
		ResponseStruct: response,
		LatencyInMs:    elapsedTime,
		Error:          err,
	}

	return clientResponse
}

// Execute performs an HTTP request using the given hystrix client and configuration.
func execute[T any](ctx context.Context, clientHttp model.ClientHttp, req model.ClientHTTPRequest) (interface{}, error) {
	var err error
	var httpRequest *http.Request
	var httpResponse *http.Response
	var responseBody []byte
	responseObject := new(T)

	if httpRequest, err = createHTTPRequest(ctx, req); err == nil {
		if httpResponse, err = clientHttp.Client.Do(httpRequest); err == nil {
			defer httpResponse.Body.Close()
			if responseBody, err = io.ReadAll(httpResponse.Body); err == nil {
				err = json.Unmarshal(responseBody, responseObject)
				if err == nil {
					return responseObject, nil
				}
			}
		}
	}

	return nil, err
}

func createHTTPRequest(ctx context.Context, req model.ClientHTTPRequest) (*http.Request, error) {
	var httpRequest *http.Request
	var bodyBytes []byte
	var err error

	if bodyBytes, err = mustMarshalBody(req.Body); err == nil {
		if httpRequest, err = http.NewRequestWithContext(ctx, req.Method, req.URL, bytes.NewBuffer(bodyBytes)); err == nil {
			for key, value := range req.Headers {
				httpRequest.Header.Add(key, value)
			}
			return httpRequest, nil
		}
	}
	return nil, err
}

func mustMarshalBody(body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch t := body.(type) {
	case string:
		return []byte(t), nil
	case []byte:
		return t, nil
	default:
		return json.Marshal(body)
	}
}
