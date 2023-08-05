package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const pathToConfig = "config.yaml"

type Config struct {
	HttpPort     string `yaml:"httpPort"`
	GrpcPort     string `yaml:"grpcPort"`
	MigrationDir string `yaml:"migrationDir"`
	DBConString  string `yaml:"dbConString"`
	Kafka        struct {
		Brokers []string `yaml:"brokers"`
		Topic   string   `yaml:"topic"`
	} `yaml:"kafka"`
	Name string `yaml:"name"`
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
