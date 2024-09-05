package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Logger LoggerConf `yaml:"logger"`
	JWT    JWTConf    `yaml:"jwt"`
	Http   HttpConf   `yaml:"http"`
}

type LoggerConf struct {
	Level string `yaml:"level" env-default:"INFO"`
}

type JWTConf struct {
	AccessTokenExpired string `yaml:"accessTokenExpired"`
}

type HttpConf struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

func NewConfig() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("Path config is empty")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("Config not found")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("Failed to read config: " + err.Error())
	}
	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
