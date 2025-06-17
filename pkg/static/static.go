package static

const (
	EnvPrefix             = "DCV"
	EnvAppPrefix          = EnvPrefix + "_APP_PREFIX"
	EnvServiceEnvironment = EnvPrefix + "_ENV"
)

const (
	ServiceEnvironmentLocal = "local"
)

const (
	EnvDbDriver   = EnvPrefix + "_DB_DRIVER"
	EnvDbDsn      = EnvPrefix + "_DB_DSN"
	EnvRouter     = EnvPrefix + "_ROUTER"
	EnvServerPort = EnvPrefix + "_SERVER_PORT"
	EnvJwtSecret  = EnvPrefix + "_JWT_SECRET"
)

const (
	DatabaseDriverPostgres  = "postgres"
	DatabaseDriverMysql     = "mysql"
	DatabaseDriverSqlLite   = "sqlite"
	DatabaseDriverSqlserver = "sqlserver"
)

const (
	HandlerUser = "user"
)
