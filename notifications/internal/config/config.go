package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const pathToConfig = "config.yaml"

type Config struct {
	HttpPort string `yaml:"httpPort"`
	GrpcPort string `yaml:"grpcPort"`
	Telegram struct {
		Token  string `yaml:"token"`
		ChatID int64  `yaml:"chatID"`
	} `yaml:"telegram"`
	Kafka struct {
		Brokers []string `yaml:"brokers"`
		Topic   string   `yaml:"topic"`
	} `yaml:"kafka"`
	RetryCount    int    `yaml:"retryCount"`
	Name          string `yaml:"name"`
	DBConString   string `yaml:"dbConString"`
	RedisHost     string `yaml:"redisHost"`
	RedisPassword string `yaml:"redisPassword"`
}

var AppConfig = Config{}

func Init() error {
	rawYaml, err := os.ReadFile(pathToConfig)
	if err != nil {
		return fmt.Errorf("read config file: %w", err)
	}

	err = yaml.Unmarshal(rawYaml, &AppConfig)
	if err != nil {
		return fmt.Errorf("parse config file: %w", err)
	}

	return nil
}
