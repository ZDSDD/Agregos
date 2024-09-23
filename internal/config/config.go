package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (*Config, error) {
	var cfg = &Config{}
	cfgDir, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}
	plan, _ := os.ReadFile(cfgDir)
	err = json.Unmarshal(plan, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", homeDir, configFileName), nil
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName

	marshalledConfig, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	cfgPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	if err = os.WriteFile(cfgPath, marshalledConfig, 2); err != nil {
		return err
	}
	return nil
}
