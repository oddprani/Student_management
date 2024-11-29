package config

import (
	"os"

	"github.com/joho/godotenv"
)

var ServerConfig *Config

type Config struct {
	DataSourceName string
	Port           string
	Environment    string
	SecretKey      string
	Debug          bool
}

func Load() (*Config, error) {
	// Load .env file if it exists
	godotenv.Load()

	config := &Config{}

	// Database
	config.DataSourceName = getEnv("DSN", "db.sqlite3")

	// Server
	config.Port = getEnv("PORT", "3000")
	config.Environment = getEnv("ENVIRONMENT", "development")
	config.Debug = config.Environment == "development"
	config.SecretKey = getEnv("SECRET_KEY", "secret")

	ServerConfig = config

	return config, nil
}

func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
