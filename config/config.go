package config

import (
	"encoding/json"
	"os"

	"github.com/go-pg/pg"
)

// App's config
type config struct {
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

	return appConfig
}

func connect() *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     Config.Database.Addr,
		User:     Config.Database.Username,
		Password: Config.Database.Password,
		Database: Config.Database.DatabaseName,
	})
}

var Config = loadConfig()
var DB = connect()