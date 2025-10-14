package config

import (
	"os"
	"strconv"
	"strings"
)

type ServerConfig struct {
	Port         string
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
	CORSOrigins  []string
}

func LoadDatabaseURL() (string, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:yourpassword@localhost:5432/echo_db?sslmode=disable"
	}

	return dsn, nil
}

func LoadServerConfig() (*ServerConfig, error) {
	cfg := &ServerConfig{
		Port:         getEnv("SERVER_PORT", "8080"),
		ReadTimeout:  getEnvAsInt("SERVER_READ_TIMEOUT", 30),
		WriteTimeout: getEnvAsInt("SERVER_WRITE_TIMEOUT", 30),
		IdleTimeout:  getEnvAsInt("SERVER_IDLE_TIMEOUT", 60),
		CORSOrigins:  getEnvAsSlice("SERVER_CORS_ORIGINS", []string{"*"}),
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func getEnvAsSlice(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		return strings.Split(value, ",")
	}
	return defaultValue
}
