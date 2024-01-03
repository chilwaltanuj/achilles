package model

import (
	"context"
)

type Client interface {
	Execute(context.Context, ClientHTTPRequest, interface{}) (any, error)
}

type ClientHTTPRequest struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    interface{}
}

type ClientResponseDetails struct {
	IsSuccessful   bool
	Status         int
	ResponseStruct interface{}
	LatencyInMs    float32
	Error          error
}
