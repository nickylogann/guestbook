package config

// databaseConfig stores database configurations
type databaseConfig struct {
	Postgres postgresConfig `toml:"Postgres"`
	Redis    redisConfig    `toml:"Redis"`
}

// postgresConfig abstracts a postgres connection config
type postgresConfig struct {
	Name     string
	Host     string
	Password string
	User     string
}

// redisConfig abstracts a redis connection config
type redisConfig struct {
	Connection string
	MaxActive  int
}
