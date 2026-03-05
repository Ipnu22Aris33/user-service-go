package config

type JWTConfig struct {
	Secret  string
	Expired string
}

func LoadJWTConfig() JWTConfig {

	return JWTConfig{
		Secret:  getEnv("JWT_SECRET", "secret"),
		Expired: getEnv("JWT_EXPIRED", "24h"),
	}
}