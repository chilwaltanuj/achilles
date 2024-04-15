package model

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type Client interface {
	Execute(context.Context, ClientHTTPRequest, interface{}) (any, error)
}
type ClientHttp struct {
	Client *resty.Client
	Logger *logrus.Logger
}

type ClientHTTPRequest struct {
	Method         string
	URL            string
	HystrixCommand string
	Headers        map[string]string
	Body           interface{}
}

type ClientResponseDetails struct {
	IsSuccessful   bool
	Status         int
	ResponseStruct interface{}
	LatencyInMs    float32
	Error          error
}
