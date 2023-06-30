package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Infra Infra
}

type Infra struct {
	Postgres Postgres
}

type Postgres struct {
	DatabaseName          string `envconfig:"DATABASE_NAME" default:"go-challenge"`
	User                  string `envconfig:"DATABASE_USER" default:"postgresql"`
	Password              string `envconfig:"DATABASE_PASSWORD" default:"postgresql"`
	Host                  string `envconfig:"DATABASE_HOST_DIRECT" default:"localhost"`
	Port                  string `envconfig:"DATABASE_PORT_DIRECT" default:"5432"`
	PoolMinSize           int32  `envconfig:"DATABASE_POOL_MIN_SIZE" default:"2"`
	PoolMaxSize           int32  `envconfig:"DATABASE_POOL_MAX_SIZE" default:"10"`
	PoolMaxConnLifetime   string `envconfig:"DATABASE_POOL_MAX_CONN_LIFETIME"`
	PoolMaxConnIdleTime   string `envconfig:"DATABASE_POOL_MAX_CONN_IDLE_TIME"`
	PoolHealthCheckPeriod string `envconfig:"DATABASE_POOL_HEALTHCHECK_PERIOD"`
	SSLMode               string `envconfig:"DATABASE_SSLMODE" default:"disable"`
	SSLRootCert           string `envconfig:"DATABASE_SSL_ROOTCERT"`
	SSLCert               string `envconfig:"DATABASE_SSL_CERT"`
	SSLKey                string `envconfig:"DATABASE_SSL_KEY"`
	Hostname              string `envconfig:"HOSTNAME"`
}

func (p Postgres) Address() string {
	if p.SSLMode == "" {
		p.SSLMode = "disable"
	}

	address := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		p.User, p.Password, p.Host, p.Port, p.DatabaseName, p.SSLMode)

	return address
}

func LoadConfig() (Config, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return Config{}, err
	}
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
