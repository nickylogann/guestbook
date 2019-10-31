package config

// databaseConfig stores database configurations
type databaseConfig struct {
	Postgres PostgresConfig `toml:"Postgres"`
	Redis    RedisConfig    `toml:"Redis"`
}

// PostgresConfig abstracts a postgres connection config
type PostgresConfig struct {
	Name     string
	Host     string
	Password string
	User     string
}

// RedisConfig abstracts a redis connection config
type RedisConfig struct {
	Connection string
	MaxActive  int
}
