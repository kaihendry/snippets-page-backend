package config

import (
	"encoding/json"
	"fmt"
	"os"
)

//Config - main app
type Config struct {
	Server struct {
		Port int `json:"port"`
	}
	Db struct {
		Address  string `json:"address"`
		Port     int    `json:"port"`
		Database string `json:"database"`
		User     string `json:"user"`
		Password string `json:"password"`
	}
	Logger struct {
		Format string `json:"format"`
	}
}

//Load - load from json
func Load(filePath string) Config {
	config := Config{}
	file, error := os.Open(filePath)
	if error != nil {
		fmt.Println(error)
	}
	defer file.Close()
	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&config)
	return config
}
