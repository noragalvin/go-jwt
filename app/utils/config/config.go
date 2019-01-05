package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Database Database
}

var config *Config

type Database struct {
	DBName     string
	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
}

func InitConfig() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	configuration := &Config{}
	error := json.Unmarshal(data, configuration)

	if error != nil {
		panic(error)
	}
	config = configuration
}

func Get() *Config {
	return config
}
