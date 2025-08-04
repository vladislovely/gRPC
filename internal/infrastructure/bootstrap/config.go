package bootstrap

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBUser string `envconfig:"DB_USER" required:"true"`
	DBPass string `envconfig:"DB_PASS" required:"true"`
	DBHost string `envconfig:"DB_HOST" required:"true"`
	DBPort int    `envconfig:"DB_PORT" required:"true"`
	DBName string `envconfig:"DB_NAME" required:"true"`
}

func ConfigFromEnv() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
