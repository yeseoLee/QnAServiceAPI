package main

import (
	"qna/config"
	"qna/datasource"
	"qna/question"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// datasource
	var ds datasource.DataSource
	if config.Conf.DB.DBType == "mysql" {
		ds = datasource.
			MySQLInstance().
			MySQLConnectionInfo(
				config.Conf.DB.Mysql.Host,
				config.Conf.DB.Mysql.Port,
				config.Conf.DB.Mysql.DatabaseName,
				config.Conf.DB.Mysql.User,
				config.Conf.DB.Mysql.Password).
			MySQLConnect()
	} else if config.Conf.DB.DBType == "sqlie3" {
		ds = datasource.
			SqliteInstance().
			SqliteConnectionInfo("").
			SqliteConnect()
	}

	////// Echo Web Framework //////
	e := echo.New()
	// Middleware
	registMiddleware(e)
	// Routes
	registRoutes(ds, e)
	// Start server
	runServer(e)
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
	e.Logger.Fatal(e.Start(config.Conf.Service.Port))
}
