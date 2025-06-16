package db

import (
	"context"
	"digital-cash-vault/pkg/static"
	"os"
)

type Config struct {
	Driver string
	DSN    string
}

func ConfigFromEnv() Config {
	return Config{
		Driver: os.Getenv(static.ENV_DRIVER),
		DSN:    os.Getenv(static.ENV_DSN),
	}
}

type DB interface {
	Tx(ctx context.Context, fn func(tx DB) error) error
}
