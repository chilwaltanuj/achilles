package config

import (
	"achilles/constant"
	"achilles/model"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// PrepareAndFetchApplicationConfiguration loads the application configuration.
// It first consults the oracle (config.default.json) and then seeks additional wisdom (environment-specific settings).
func PrepareAndFetchApplicationConfiguration() (*model.ApplicationConfiguration, error) {
	viper.SetConfigName("config.default")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config") // Path to the main config file

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf(constant.ConfigDefaultReadError, err)
	}
	log.Println(constant.ConfigDefaultLoadSuccess)

	environment := viper.GetString("deployment_environment")
	envConfigFile := fmt.Sprintf("config.%s.json", environment)

	viper.SetConfigName(envConfigFile)
	if err := viper.MergeInConfig(); err != nil {
		log.Printf(constant.ConfigEnvLoadFailure, environment, envConfigFile, err)
	} else {
		log.Printf(constant.ConfigEnvLoadSuccess, envConfigFile)
	}

	var applicationConfiguration model.ApplicationConfiguration
	if err := viper.Unmarshal(&applicationConfiguration); err != nil {
		return nil, fmt.Errorf(constant.ConfigUnmarshalError, err)
	}

	log.Println(constant.ConfigReady)
	return &applicationConfiguration, nil
}
