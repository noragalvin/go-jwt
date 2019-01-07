package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Domain        string `json:"domain"`
	SessionSecret string `json:"session_secret"`
	Port          int    `json:"port"`
	Database      DatabaseConfig
}

var config *Config

type DatabaseConfig struct {
	Name     string
	Host     string
	Port     string
	Username string
	Password string
}

func init() {
	data, err := ioutil.ReadFile("config.json")
	// log.Println(data)
	if err != nil {
		panic(err)
	}

	configuration := &Config{}
	error := json.Unmarshal(data, &configuration)

	if error != nil {
		panic(error)
	}
	config = configuration
}

func Get() *Config {
	return config
}
