package config

import (
	"encoding/json"
	"os"
)

// App's config
var Config *config

func init() {
	Config = loadConfig()
}

type config struct {
	Port     string `json:"port"`
	Database struct {
		Addr         string `json:"addr"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		DatabaseName string `json:"database_name"`
	} `json:"database"`
}

// LoadConfig load the config.json file and return its data
func loadConfig() *config {
	appConfig := &config{}

	configFile, err := os.Open("config/config.json")
	defer configFile.Close()
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(configFile).Decode(&appConfig)
	if err != nil {
		panic(err)
	}

	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		appConfig.Port = envPort
	}

	return appConfig
}
