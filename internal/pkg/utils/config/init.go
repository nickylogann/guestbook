package config

import "os"

var cfg *RuntimeConfig

// RuntimeConfig stores the global app configuration
type RuntimeConfig struct {
	Database databaseConfig `toml:"Database"`
}

// NewConfig initializes the global config
func NewConfig() RuntimeConfig {
	if cfg != nil {
		return *cfg
	}

	cfg = loadConfigFromEnv()
	return *cfg
}

func loadConfigFromEnv() *RuntimeConfig {
	return &RuntimeConfig{
		Database: databaseConfig{
			Postgres: PostgresConfig{
				Host:     os.Getenv("PG_HOST"),
				Name:     os.Getenv("PG_NAME"),
				User:     os.Getenv("PG_USER"),
				Password: os.Getenv("PG_PASS"),
			},
			Redis: RedisConfig{},
		},
	}
}
