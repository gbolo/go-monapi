package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"strings"
)

func loadConfig() {

	// Set the config options
	viper.SetConfigType("yaml")
	viper.SetConfigName(configName)

	viper.AddConfigPath("./")
	var alternativeCfgPath = os.Getenv("MONAPI_CFG_PATH")
	if alternativeCfgPath != "" {
		viper.AddConfigPath(alternativeCfgPath)
	}

	// Read in environment variables that match
	viper.SetEnvPrefix(envVarPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		logger.Warning("No config file loaded. Using defaults")
	} else {
		logger.Info("Loaded config file:", viper.ConfigFileUsed())
	}

}

func getServerConfig() string {

	viper.SetDefault("server.bind_address", "0.0.0.0")
	viper.SetDefault("server.bind_port", "8080")

	return fmt.Sprintf(
		"%s:%s",
		viper.GetString("server.bind_address"),
		viper.GetString("server.bind_port"),
	)
}
