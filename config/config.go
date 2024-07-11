package config

import (
	"os"
)

type Config struct {
	MongoURI string
	DBName   string
}

func LoadConfig() *Config {
	return &Config{
		MongoURI: getEnv("MONGO_URI", "mongodb://mongo:27017"), // Убедитесь, что это правильный URI
		DBName:   getEnv("DB_NAME", "your_db"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
