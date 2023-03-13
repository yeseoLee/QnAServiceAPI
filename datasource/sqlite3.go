package datasource

import (
	"database/sql"
	"fmt"
	"log"

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
	s.createTable()
	return s
}

func (s sqlite) GetConnection() (*sql.DB, error) {
	if !s.ok {
		return nil, fmt.Errorf("fail to connect sqlite db")
	}
	return s.dbpool, nil
}

func (s sqlite) createTable() {
	createQuestionTableQuery := `
		CREATE TABLE IF NOT EXISTS tbQuestion ( 
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		writerId VARCHAR(255) NOT NULL,
		title VARCHAR(255) NOT NULL,
		content TEXT,
		tags TEXT,
		images TEXT,
		isAccept BOOLEAN NOT NULL CHECK (isAccept IN (0, 1)) DEFAULT 0,
		createdAt DATETIME,
		updatedAt DATETIME,
		deletedAt DATETIME
		);
	`
	createAnswerTableQuery := `
		CREATE TABLE IF NOT EXISTS tbAnswer ( 
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		questionId INTEGER UNIQUE NOT NULL,
		writerId VARCHAR(255) NOT NULL,
		content TEXT,
		images TEXT,
		isAccept BOOLEAN NOT NULL CHECK (isAccept IN (0, 1)) DEFAULT 0,
		createdAt DATETIME,
		updatedAt DATETIME,
		deletedAt DATETIME
		);
	`
	createCommentTableQuery := `
		CREATE TABLE IF NOT EXISTS tbComment ( 
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		questionId INTEGER UNIQUE NOT NULL,
		answerId INTEGER UNIQUE NOT NULL,
		writerId VARCHAR(255) NOT NULL,
		content TEXT,
		createdAt DATETIME,
		deletedAt DATETIME
		);
	`
	if _, err := s.dbpool.Exec(createQuestionTableQuery); err != nil {
		log.Fatal(fmt.Errorf("createQuestionTableQuery error: %w", err))
	}
	if _, err := s.dbpool.Exec(createAnswerTableQuery); err != nil {
		log.Fatal(fmt.Errorf("createAnswerTableQuery error: %w", err))
	}
	if _, err := s.dbpool.Exec(createCommentTableQuery); err != nil {
		log.Fatal(fmt.Errorf("createCommentTableQuery error: %w", err))
	}
}
