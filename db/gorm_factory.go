package db

import (
	"digital-cash-vault/pkg/static"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func openGORM(cfg Config) (*gormDB, error) {
	var dial gorm.Dialector

	switch cfg.Driver {
	case static.DatabaseDriverPostgres:
		dial = postgres.Open(cfg.DSN)
	case static.DatabaseDriverMysql:
		dial = mysql.Open(cfg.DSN)
	case static.DatabaseDriverSqlLite:
		dial = sqlite.Open(cfg.DSN)
	case static.DatabaseDriverSqlserver:
		dial = sqlserver.Open(cfg.DSN)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", cfg.Driver)
	}

	db, err := gorm.Open(dial, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return &gormDB{
		db,
	}, nil
}
