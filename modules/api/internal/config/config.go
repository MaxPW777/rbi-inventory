package config

import "os"

type Config struct {
	DatabaseURL string
	Port        string
}

func Load() *Config {
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", ""),
		Port:        getEnv("PORT", "5432"),
	}
}

func getEnv(envValue string, defaultValue string) string {
	if value := os.Getenv(envValue); value != "" {
		return value
	}
	return defaultValue
}
