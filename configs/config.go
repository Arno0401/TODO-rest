package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		SSLMode  string `yaml:"ssl_mode"`
	} `yaml:"database"`
	Token struct {
		Secret  string        `yaml:"secret"`
		Access  time.Duration `yaml:"access_exp"`
		Refresh time.Duration `yaml:"refresh_exp"`
	} `yaml:"token"`
}

var DBConfig Config

func GetDBConfig() {
	file, err := os.Open("configs.yml")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&DBConfig); err != nil {
		panic(err)
	}
}
