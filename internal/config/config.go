package config

import "github.com/joho/godotenv"

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

func LoadConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		App:      LoadAppConfig(),
		Database: LoadDatabaseConfig(),
		JWT:      LoadJWTConfig(),
	}

	return cfg, nil
}