package cmd

import "os"

type Config struct {
	Environment string `arg:"env:ENVIRONMENT"`
	Port        string `arg:"env:SERVER_PORT"`
	Name        string `arg:"env:SERVER_NAME"`
	URL         string `arg:"env:DATABASE_URL"`
}

func DefaultConfig() *Config {
	return &Config{
		Environment: getEnv("ENVIRONMENT", "dev"),
		Name:        getEnv("SERVER_NAME", "backend"),
		Port:        getEnv("SERVER_PORT", "8080"),
		URL:         getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
