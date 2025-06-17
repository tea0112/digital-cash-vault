package db

import (
	"context"

	"gorm.io/gorm"
)

type gormDB struct {
	*gorm.DB
}

func (g gormDB) Tx(ctx context.Context, fn func(tx DB) error) error {
	return g.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(gormDB{tx})
	})
}
func (g gormDB) First(ctx context.Context, dest any, query string, args ...any) error {
	return g.DB.WithContext(ctx).Where(query, args...).First(dest).Error
}
func (g gormDB) Find(ctx context.Context, dest any, query string, args ...any) error {
	return g.DB.WithContext(ctx).Where(query, args...).Find(dest).Error
}

func (g gormDB) Exec(ctx context.Context, query string, args ...any) (int64, error) {
	res := g.DB.WithContext(ctx).Exec(query, args...)
	return res.RowsAffected, res.Error
}
func (g gormDB) Raw() any { return g.DB }
