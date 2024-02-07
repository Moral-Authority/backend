package cmd

import "os"

// These variables are meant to inserted as env vars
// you can inject them through orchestrator secrets or command line args or OS env vars
// Precedence: Command-line flags > Env vars > Default values.

type Config struct {
	Environment string `arg:"env:ENVIRONMENT"`
	ServerConfig
	DatabaseConfig
}

type ServerConfig struct {
	Port string `arg:"env:SERVER_PORT"`
	Name string `arg:"env:SERVER_NAME"`
}

type DatabaseConfig struct {
	DatabaseName           string `arg:"env:DATABASE_NAME"`
	DatabaseUsername       string `arg:"env:DATABASE_USERNAME"`
	DatabasePassword       string `arg:"env:DATABASE_PASSWORD"`
	DatabaseConnectionUrl  string `arg:"env:DATABASE_CONNECTION_URL"`
	DatabaseConnectionPort string `arg:"env:DATABASE_CONNECTION_PORT"`
}

func DefaultConfiguration() *Config {
	return &Config{
		Environment: os.Getenv("ENVIRONMENT"),
		ServerConfig: ServerConfig{
			Name: os.Getenv("SERVER_NAME"),
			Port: os.Getenv("SERVER_PORT"),
		},
		DatabaseConfig: DatabaseConfig{
			DatabaseName:           os.Getenv("DATABASE_NAME"),
			DatabaseUsername:       os.Getenv("DATABASE_USERNAME"),
			DatabasePassword:       os.Getenv("DATABASE_PASSWORD"),
			DatabaseConnectionUrl:  os.Getenv("DATABASE_CONNECTION_URL"),
			DatabaseConnectionPort: os.Getenv("DATABASE_CONNECTION_PORT"),
		},
	}
}
