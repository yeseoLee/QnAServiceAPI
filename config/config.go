package config

import (
	"encoding/json"
	"os"
	"sync"
)

// JSON struct
type (
	Config struct {
		Service service `json: "service"`
		Mysql   mysql   `json: "mysql"`
	}
	service struct {
		Port string `json: "port"`
		Root string `json: "root"`
	}
	mysql struct {
		Host         string `json: "host"`
		Port         int    `json: "port"`
		DatabaseName string `json: "databaseName"`
		User         string `json: "user"`
		Password     string `json: "password"`
	}
)

var once sync.Once
var instance *Config

func GetInstance() *Config {
	once.Do(func() {
		instance = loadConfigJSON("config/config.json")
	})
	return instance
}

func loadConfigJSON(path string) (c *Config) {
	file, _ := os.Open(path)
	defer file.Close()
	decoder := json.NewDecoder(file)
	decoder.Decode(&c)
	return c
}
