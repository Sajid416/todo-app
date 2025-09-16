package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	HttpPort  int
	DBUrl     string
	JWTSecret string
}

// GetConfig loads configuration from environment variables
func GetConfig() *Config {
	httpPort := getEnvAsInt("HTTP_PORT", 8080)
	dbUrl := getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/todo_app?sslmode=disable")
	jwtSecret := getEnv("JWT_SECRET", "supersecretkey")

	return &Config{
		HttpPort:  httpPort,
		DBUrl:     dbUrl,
		JWTSecret: jwtSecret,
	}
}

// Helper: Get env var with fallback
func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// Helper: Get int env var with fallback
func getEnvAsInt(key string, fallback int) int {
	if valueStr, exists := os.LookupEnv(key); exists {
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			log.Printf("Invalid value for %s: %s, using fallback %d", key, valueStr, fallback)
			return fallback
		}
		return value
	}
	return fallback
}
