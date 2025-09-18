package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	HttpPort  int
	UserDBUrl string
	TodoDBUrl string
	JWTSecret string
	JWTRefresh string
}

func GetConfig() *Config {
	httpPort := getEnvAsInt("HTTP_PORT", 8080)
	jwtSecret := getEnv("JWT_SECRET", "myverysecretkey")
	jwtRefresh:=getEnv("JWT_REFRESH","myveryrefreshkey")
    
	userDBUrl := buildDBUrl(
		getEnv("DB_USER", "user"),
		getEnv("DB_PASSWORD", "user_pass"),
		getEnv("DB_HOST", "localhost"),
		getEnvAsInt("DB_PORT", 5432),
		getEnv("DB_NAME", "user_db"), // default user DB name
	)

	ProductDBUrl := buildDBUrl(
		getEnv("DB_USER", "user"),
		getEnv("DB_PASSWORD", "user_pass"),
		getEnv("DB_HOST", "localhost"),
		getEnvAsInt("DB_PORT", 5432),
		getEnv("DB_NAME", "todo_app"),
	)

	return &Config{
		HttpPort:  httpPort,
		UserDBUrl: userDBUrl,
		TodoDBUrl: ProductDBUrl,
		JWTSecret: jwtSecret,
		JWTRefresh: jwtRefresh,
	}
}


func buildDBUrl(user, password, host string, port int, dbname string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
}

func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

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
