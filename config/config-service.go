package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Config struct {
	IsFileBasedGraphReadingEnabled bool `json:"feature.flag.read.graph.from.file.enabled"`
}

type ConfigService struct {
	PathToConfigFile string
}

var (
	instance *ConfigService
	once     sync.Once
)

func GetInstance(pathToConfigFile string) *ConfigService {
	once.Do(func() {
		instance = &ConfigService{pathToConfigFile}
		fmt.Println("Confing instance created")
	})
	return instance
}

func (configService *ConfigService) ReadConfig() Config {
	file, err := os.Open(configService.PathToConfigFile)
	if err != nil {
		fmt.Printf("Erropr during reading file.: %v\n", err)
		return Config{}
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		fmt.Printf("Error during parsing JSON: %v\n", err)
		return Config{}
	}

	return config
}
