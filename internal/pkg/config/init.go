package config

import "os"

var cfg *RuntimeConfig

// RuntimeConfig stores the global app configuration
type RuntimeConfig struct {
	Database databaseConfig `toml:"Database"`
}

// NewConfig initializes the global config
func NewConfig() *RuntimeConfig {
	if cfg != nil {
		return cfg
	}

	cfg = loadConfigFromEnv()
	return cfg
}

func loadConfigFromEnv() *RuntimeConfig {
	return &RuntimeConfig{
		Database: databaseConfig{
			Postgres: postgresConfig{
				Name:     os.Getenv("DB_NAME"),
				Host:     os.Getenv("DB_HOST"),
				Password: os.Getenv("DB_PASS"),
				User:     os.Getenv("DB_USER"),
			},
			Redis: redisConfig{},
		},
	}
}