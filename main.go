package main

import (
	"fmt"
	"log"
	"qna/config"
	"qna/datasource"
	"qna/question"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// datasource
	ds := setDataSource()

	////// Echo Web Framework //////
	e := echo.New()
	// Middleware
	registMiddleware(e)
	// Routes
	registRoutes(ds, e)
	// Start server
	runServer(e)
}

func setDataSource() (ds datasource.DataSource) {
	if config.GetInstance().DB.DBType == datasource.DRIVER_NAME_MYSQL { // mysql
		ds = datasource.
			MySQLInstance().
			MySQLConnectionInfo(
				config.GetInstance().DB.Mysql.Host,
				config.GetInstance().DB.Mysql.Port,
				config.GetInstance().DB.Mysql.DatabaseName,
				config.GetInstance().DB.Mysql.User,
				config.GetInstance().DB.Mysql.Password).
			MySQLConnect()
	} else if config.GetInstance().DB.DBType == datasource.DRIVER_NAME_SQLITE { // sqlite
		ds = datasource.
			SqliteInstance().
			SqliteConnectionInfo(config.GetInstance().DB.Sqlite.FilePath).
			SqliteConnect()
	} else {
		log.Fatal(fmt.Errorf("config.yaml DBType Error"))
	}
	return
}

func registMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func registRoutes(ds datasource.DataSource, e *echo.Echo) {
	// Question
	question.RegistQuestionRoute(ds, e)
}

func runServer(e *echo.Echo) {
	e.Logger.Fatal(e.Start(config.GetInstance().Service.Port))
}
