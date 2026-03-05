package config

import "os"

type AppConfig struct {
	Name string
	Port string
	Env  string
}

func LoadAppConfig() AppConfig {
	return AppConfig{
		Name: getEnv("APP_NAME", "user-service"),
		Port: getEnv("APP_PORT", "8080"),
		Env:  getEnv("APP_ENV", "development"),
	}
}

func getEnv(key string, fallback string) string {

	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}