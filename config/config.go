package config

import (
	"encoding/json"
	"os"
)

//Config - main app
type Config struct {
	App struct {
		Frontend string `json:"frontend"`
		Backend  string `json:"backend"`
		Port     string `json:"port"`
	}
	Db struct {
		ConnectionAddress string `json:"connectionAddress"`
		Database          string `json:"database"`
	}
	JWT struct {
		Secret string `json:"secret"`
	}
	CORS struct {
	}
	TLS struct {
		Enable bool   `json:"enable"`
		Cert   string `json:"cert"`
		Key    string `json:"key"`
	}
}

//Load - load from json
func Load(filePath string) (Config, error) {
	config := Config{}
	file, error := os.Open(filePath)
	defer file.Close()
	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&config)
	return config, error
}
