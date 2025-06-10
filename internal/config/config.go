// Package config provides configuration loading from YAML files
package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// ServerConfig holds server-related settings
type ServerConfig struct {
	Port int `yaml:"port"`
}

// LoggingConfig holds logging-related settings
type LoggingConfig struct {
	Path string `yaml:"path"`
}

// Config aggregates all service configurations
type Config struct {
	Server  ServerConfig  `yaml:"server"`
	Logging LoggingConfig `yaml:"logging"`
}

// LoadConfig loads the configuration from the given YAML file path
func LoadConfig(path string) (*Config, error) {
	// #nosec G304 -- config path is trusted and not user-controlled
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
