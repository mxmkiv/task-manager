package main

import (
	"task-manager/internal/routes"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {

	e := echo.New()

	// log all requests
	e.Use(middleware.RequestLogger())

	routes.Setup(e)

	err := e.Start(":8888")
	if err != nil {
		e.Logger.Error("server start error", "error", err)
	}
}
