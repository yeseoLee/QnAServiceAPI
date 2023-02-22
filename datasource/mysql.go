package datasource

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	driverName     string
	dataSourceName string
	dbpool         *sql.DB
	ok             bool
}

var mysqlInstance mysql

func MySQLInstance() mysql {
	return mysqlInstance
}

func (m mysql) MySQLConnectionInfo(ip string, port int, dbName string, id string, pw string) mysql {
	m.driverName = DRIVER_NAME_MYSQL
	m.dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", id, pw, ip, port, dbName)
	return m
}

func (m mysql) MySQLConnect() mysql {
	db, err := sql.Open(m.driverName, m.dataSourceName)
	if err != nil {
		m.ok = false
	}
	m.dbpool = db
	m.ok = true
	return m
}

func (m mysql) GetConnection() (*sql.DB, error) {
	if !m.ok {
		return nil, fmt.Errorf("fail to connect mysql db")
	}
	return m.dbpool, nil
}
