package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const pathToConfig = "config.yaml"

type Loms struct {
	Addr string `yaml:"addr"`
}

type Products struct {
	Addr         string `yaml:"addr"`
	Token        string `yaml:"token"`
	RateLimit    int    `yaml:"rateLimit"`
	WorkersCount int    `yaml:"workersCount"`
}

type Config struct {
	HttpPort        string   `yaml:"httpPort"`
	GrpcPort        string   `yaml:"grpcPort"`
	MigrationDir    string   `yaml:"migrationDir"`
	DBConString     string   `yaml:"dbConString"`
	LomsService     Loms     `yaml:"loms"`
	ProductsService Products `yaml:"products"`
	Name            string   `yaml:"name"`
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
