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
			DatabaseName:          os.Getenv("DATABASE_NAME"),
			DatabaseUsername:      os.Getenv("DATABASE_USERNAME"),
			DatabasePassword:      os.Getenv("DATABASE_PASSWORD"),
			DatabaseConnectionUrl: os.Getenv("DATABASE_URL"),
			//DatabaseConnectionPort: os.Getenv("DATABASE_CONNECTION_PORT"),
		},
	}
}

//// DefaultConfiguration return dev configuration as default config
//func DefaultConfiguration() *Config {
//	return &Config{
//		Environment: "dev",
//		ServerConfig: ServerConfig{
//			Name: "backend",
//			Port: "8080",
//		},
//		DatabaseConfig: DatabaseConfig{
//			DatabaseName:          "postgres",
//			DatabaseUsername:      "postgres",
//			DatabasePassword:      "postgres",
//			DatabaseConnectionUrl: "postgres://ehxxgxaamkwsaz:15acfc1d3ddc307e898673d17a669f2169913ef6467380439410f045e42f3e4d@ec2-52-0-79-72.compute-1.amazonaws.com:5432/d7mdhbr0m054dr",
//			//DatabaseConnectionPort: "5432",
//		},
//	}
//}
