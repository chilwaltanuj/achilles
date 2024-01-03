package model

import (
	"github.com/gojek/heimdall/v7/hystrix"
	"github.com/sirupsen/logrus"
)

type ClientHttp struct {
	Client *hystrix.Client
	Logger *logrus.Logger
}

type Joke_Random struct {
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
	ID        int    `json:"id"`
}
