package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HttpPort     string  `yaml:"http_port" env-default:"9000"`
	BalancerName string  `yaml:"balancer_name" env-default:"round_robin"`
	MatrixSize   int     `yaml:"matrix_size" env-default:"150"`
	TimeSleep    float64 `yaml:"time_sleep" env-default:"5"`
}

func New() (Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return Config{}, fmt.Errorf("CONFIG_PATH is not set")
	}

	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return Config{}, fmt.Errorf("config file %s not found", configPath)
	}

	var cfg Config

	err = cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("cleanenv.ReadConfig: %w", err)
	}

	return cfg, nil
}
