package config

import (
	"log"

	"github.com/spf13/viper"
)

// DatabaseConfig struct holds the configuration for the database connection
type DatabaseConfig struct {
	DSN string
}

// Config struct holds the application configuration
type Config struct {
	DATABASE_DEV_URL   string `mapstructure:"DATABASE_DEV_URL"`
	DATABASE_LOCAL_URL string `mapstructure:"DATABASE_LOCAL_URL"`
}

// LoadConfig loads configuration from environment variables or a file
func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

// GetDatabaseConfig returns the database configuration based on the environment
func GetDatabaseConfig(env string) DatabaseConfig {
	var dsn string
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	// Determine the DSN based on the environment
	switch env {
	case "development":
		dsn = cfg.DATABASE_DEV_URL
	case "local":
		dsn = cfg.DATABASE_LOCAL_URL
	default:
		log.Fatalf("Unknown environment: %s", env)
	}

	return DatabaseConfig{
		DSN: dsn,
	}
}
