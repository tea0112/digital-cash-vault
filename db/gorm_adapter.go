package db

import (
	"context"
	"digital-cash-vault/internal/infrastructure/query"

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
	db := g.DB.WithContext(ctx)
	if query == "" {
		return db.Find(dest).Error
	}

	return db.Where(query, args...).Find(dest).Error
}

func (g gormDB) Exec(ctx context.Context, query string, args ...any) (int64, error) {
	res := g.DB.WithContext(ctx).Exec(query, args...)
	return res.RowsAffected, res.Error
}
func (g gormDB) Raw() any { return g.DB }

type GormQueryBuilder struct {
	db *gorm.DB
}

func NewGormQueryBuilder(db *gorm.DB) query.Builder {
	return &GormQueryBuilder{db: db}
}

func (q *GormQueryBuilder) Where(query string, args ...any) query.Builder {
	q.db = q.db.Where(query, args...)
	return q
}
func (q *GormQueryBuilder) OrderBy(orderBy string) query.Builder {
	q.db = q.db.Order(orderBy)
	return q
}
func (q *GormQueryBuilder) Limit(n int) query.Builder {
	q.db = q.db.Limit(n)
	return q
}
func (q *GormQueryBuilder) Offset(n int) query.Builder {
	q.db = q.db.Offset(n)
	return q
}
func (q *GormQueryBuilder) Preload(field string) query.Builder {
	q.db = q.db.Preload(field)
	return q
}
func (q *GormQueryBuilder) BuildNative() any {
	return q.db
}
