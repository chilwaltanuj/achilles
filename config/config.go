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
	viper.SetConfigName("config")   // Set the name of the configuration file (without extension)
	viper.SetConfigType("json")     // Set the configuration file type
	viper.AddConfigPath("./config") // Add the path to the config directory

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %s", err)
	}

	var applicationConfiguration model.ApplicationConfiguration
	if err := viper.Unmarshal(&applicationConfiguration); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %s", err.Error())
	}
	return &applicationConfiguration, nil
}
