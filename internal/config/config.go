package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func (c Config) SetUser(userName string) error {
	config, err := Read()
	if err != nil {
		fmt.Println("error reading config file")
		return err
	}
	config.CurrentUserName = userName
	return write(config)
}

func write(cfg *Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	MarshalledConfig, err := json.Marshal(cfg)
	if err != nil {
		fmt.Println("error marshalling config file")
		return err
	}
	return os.WriteFile(configFilePath, MarshalledConfig, 0644)
}

func Read() (*Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		fmt.Println("error decoding config file")
		return nil, err
	}

	return &config, nil
}

func getConfigFilePath() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path + "/" + configFileName, nil
}
