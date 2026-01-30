package routes

import (
	"task-manager/internal/handler"
	"task-manager/internal/middleware"

	"github.com/labstack/echo/v5"
)

func Setup(e *echo.Echo) {

	//api := e.Group("/api/v1")

	e.GET("/user", handler.UserHandler, middleware.TestMW)
	e.GET("/task", handler.TaskHandler)
	e.GET("/", handler.MainHandler)

}
