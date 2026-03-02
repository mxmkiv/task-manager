package routes

import (
	"task-manager/internal/handler"
	"task-manager/internal/middleware"

	"github.com/labstack/echo/v5"
)

func Setup(e *echo.Echo) {

	//api := e.Group("/api/v1")

	//auth
	e.POST("/auth/register", handler.RegisterHandler)
	//e.POST("/auth/login", )

	//protect
	e.GET("/test", handler.TestHandler, middleware.AuthWiddleware)
}
