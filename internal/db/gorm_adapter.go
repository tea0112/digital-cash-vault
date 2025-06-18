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
func (g gormDB) First(dest any, q string, args ...any) error {
	return g.DB.Raw(q, args...).First(dest).Error
}
func (g gormDB) Find(dest any, q string, args ...any) error {
	return g.DB.Raw(q, args...).Find(dest).Error
}

func (g gormDB) Exec(q string, args ...any) (int64, error) {
	res := g.DB.Exec(q, args...)
	return res.RowsAffected, res.Error
}
func (g gormDB) Raw() any { return g.DB } // gives you *gorm.DB if you really need it
