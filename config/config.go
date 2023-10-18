package config

import (
	"achilles/model"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func BuildAndGetApplicationConfiguration() (*model.ApplicationConfiguration, error) {
	gin.SetMode(gin.DebugMode)

	// Initialize Viper
	viper.SetConfigFile("./config/config.json")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %s", err)
	}

	var applicationConfiguration model.ApplicationConfiguration
	if err := viper.Unmarshal(&applicationConfiguration); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %s", err.Error())
	}
	return &applicationConfiguration, nil
}
