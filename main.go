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
	ds := datasource.
		MySQLInstance().
		MySQLConnectionInfo(
			config.GetInstance().Mysql.Host,
			config.GetInstance().Mysql.Port,
			config.GetInstance().Mysql.DatabaseName,
			config.GetInstance().Mysql.User,
			config.GetInstance().Mysql.Password).
		MySQLConnect()

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
	e.Logger.Fatal(e.Start(config.GetInstance().Service.Port))
}
