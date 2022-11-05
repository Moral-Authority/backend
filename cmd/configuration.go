package cmd

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

// DefaultConfiguration return dev configuration as default config
func DefaultConfiguration() *Config {
	return &Config{
		Environment: "dev",
		ServerConfig: ServerConfig{
			Name: "backend",
			Port: "8080",
		},
		DatabaseConfig: DatabaseConfig{
			DatabaseName:           "backend_db",
			DatabaseUsername:       "postgres",
			DatabasePassword:       "postgres",
			DatabaseConnectionUrl:  "localhost",
			DatabaseConnectionPort: "5432",
		},
	}
}
