package config

import (
	"achilles/model"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func BuildAndGetApplicationConfiguration() (*model.ApplicationConfiguration, error) {
	gin.SetMode(gin.DebugMode)

	viper.SetConfigFile("./config/config.json")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var applicationConfiguration model.ApplicationConfiguration
	if err := viper.Unmarshal(&applicationConfiguration); err != nil {
		return nil, err
	}
	return &applicationConfiguration, nil
}
