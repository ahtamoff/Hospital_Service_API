package config

import (
	"os"
)

type Config struct {
	MongoURI   string
	DBName     string
	ServerPort string
	LogSpeed   string
}

func LoadConfig() *Config {
	return &Config{
		MongoURI:   getEnv("MONGO_URI", "mongodb://mongo:27017"),
		DBName:     getEnv("DB_NAME", "hospital"),
		ServerPort: getEnv("PORT", ":8080"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
