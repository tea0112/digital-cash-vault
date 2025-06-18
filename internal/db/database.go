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
		Driver: os.Getenv(static.EnvDcvDbDriver),
		DSN:    os.Getenv(static.EnvDcvDbDsn),
	}
}

type DB interface {
	Tx(ctx context.Context, fn func(tx DB) error) error
	First(dest any, query string, args ...any) error
	Find(dest any, query string, args ...any) error
	Exec(query string, args ...any) (int64, error)
	Raw() any
}

func Open(cfg Config) (DB, error) {
	gDB, err := openGORM(cfg) // the helper you already wrote
	if err != nil {
		return nil, err
	}

	return gDB, nil
}
