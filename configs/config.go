package config

import (
	"os"
	"strconv"
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
	DBConfig.Server.Host = getEnv("SERVER_HOST", "127.0.0.1")
	DBConfig.Server.Port = getEnv("SERVER_PORT", "2211")
	DBConfig.Database.User = getEnv("DB_USER", "arno")
	DBConfig.Database.Password = getEnv("DB_PASSWORD", "5O4QjCtgfT0W71p4ugmtwb1WbwjMS2Ds")
	DBConfig.Database.DBName = getEnv("DB_NAME", "todo_px3l")
	DBConfig.Database.Host = getEnv("DB_HOST", "dpg-cubld0ogph6c73a7dls0-a")
	DBConfig.Database.Port = getEnvInt("DB_PORT", 5432)
	DBConfig.Database.SSLMode = getEnv("DB_SSL_MODE", "disable")
	DBConfig.Token.Secret = getEnv("TOKEN_SECRET", "arno")
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	}
	return defaultValue
}
