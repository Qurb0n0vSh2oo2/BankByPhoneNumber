package config

type Config struct {
	PostgresURL         string
	ComissionOfTransfer float64
}

func NewConfig() *Config {
	return &Config{
		PostgresURL:         "postgres://postgres:3302@localhost:5432/sql_bankcli",
		ComissionOfTransfer: 5.0,
	}
}
