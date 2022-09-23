package server

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func NewConfig(path string) (*Config, error) {
	buf, err := os.ReadFile("config.yml")
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err = yaml.Unmarshal(buf, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
