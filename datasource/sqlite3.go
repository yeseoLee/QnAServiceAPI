package datasource

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type sqlite struct {
	driverName     string
	dataSourceName string
	dbpool         *sql.DB
	ok             bool
}

var sqliteInstance sqlite

func SqliteInstance() sqlite {
	return sqliteInstance
}

func (s sqlite) SqliteConnectionInfo(file string) sqlite {
	s.driverName = DRIVER_NAME_SQLITE
	s.dataSourceName = file
	return s
}

func (s sqlite) SqliteConnect() sqlite {
	db, err := sql.Open(s.driverName, s.dataSourceName)
	if err != nil {
		s.ok = false
	}
	s.dbpool = db
	s.ok = true
	return s
}

func (s sqlite) GetConnection() (*sql.DB, error) {
	if !s.ok {
		return nil, fmt.Errorf("fail to connect sqlite db")
	}
	return s.dbpool, nil
}
