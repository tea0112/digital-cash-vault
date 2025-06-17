package query

type Builder interface {
	Where(query string, args ...any) Builder
	OrderBy(orderBy string) Builder
	Limit(n int) Builder
	Offset(n int) Builder
	Preload(field string) Builder
	BuildNative() any
}
