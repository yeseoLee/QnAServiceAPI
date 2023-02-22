package datasource

import "database/sql"

const (
	DRIVER_NAME_MYSQL      = "mysql"
	DRIVER_NAME_MARIADB    = "mysql"
	DRIVER_NAME_SQLITE     = "sqlite3"
	DRIVER_NAME_POSTGRESQL = "postgres"
)

type DataSource interface {
	GetConnection() (*sql.DB, error)
}
