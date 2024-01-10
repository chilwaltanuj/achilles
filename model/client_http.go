package model

import (
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type ClientHttp struct {
	Client *resty.Client
	Logger *logrus.Logger
}

type Joke_Random struct {
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
	ID        int    `json:"id"`
}
