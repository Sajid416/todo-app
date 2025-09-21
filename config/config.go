package config

import (
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type Config struct {
	HttpPort   int
	DBUrl      string
	JWTSecret  string
	JWTRefresh string
}

func GetConfig() *Config {
	httpPort := getEnvAsInt("HTTP_PORT", 8080)
	jwtSecret := getEnv("JWT_SECRET", "myverysecretkey")
	jwtRefresh := getEnv("JWT_REFRESH", "myveryrefreshkey")

	// DBUrl := buildDBUrl(
	// 	getEnv("DB_USER", "new_user"),
	// 	getEnv("DB_PASSWORD", "new_pass123"),
	// 	getEnv("DB_HOST", "localhost"),
	// 	getEnvAsInt("DB_PORT", 5432),
	// 	getEnv("DB_NAME", "user_product"), // default user DB name
	// )
	DBUrl := "postgres://new_user:new_pass123@localhost:5432/user_product?sslmode=disable"

	return &Config{
		HttpPort:   httpPort,
		DBUrl:      DBUrl,
		JWTSecret:  jwtSecret,
		JWTRefresh: jwtRefresh,
	}
}

// func buildDBUrl(user, password, host string, port int, dbname string) string {
// 	//return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
// 	return
// }

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
