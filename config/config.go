package config

import (
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
		return nil, fmt.Errorf("achilles stumbled while reading the oracle's scroll (config.default.json): %w", err)
	}
	log.Println("Achilles nods wisely, having read the oracle's scroll (config.default.json)")

	environment := viper.GetString("deployment_environment")
	envConfigFile := fmt.Sprintf("config.%s.json", environment)

	viper.SetConfigName(envConfigFile)
	if err := viper.MergeInConfig(); err != nil {
		log.Printf("Achilles raises an eyebrow: 'Are we sure about environment = %s?' Couldn't find the scroll %s: %s. Sticking to the oracle's advice (config.default.json).",
			environment, envConfigFile, err)
	} else {
		log.Printf("Achilles finds and unrolls the scroll %s, blending its secrets with the oracle's wisdom (config.default.json).", envConfigFile)
	}

	var applicationConfiguration model.ApplicationConfiguration
	if err := viper.Unmarshal(&applicationConfiguration); err != nil {
		return nil, fmt.Errorf("achilles grunts in frustration: Something's amiss with marshalling the scrolls into formation: %w", err)
	}

	log.Println("Achilles stands ready, configurations in hand and a strategy in mind.")
	return &applicationConfiguration, nil
}
