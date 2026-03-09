package routes

import (
	"task-manager/internal/handler"
	"task-manager/internal/middleware"

	"github.com/labstack/echo/v5"
)

func Setup(e *echo.Echo, h *handler.Handler) {

	//api := e.Group("/api/v1")

	//auth
	e.POST("/auth/register", h.RegisterHandler, middleware.GuestOnly)
	e.POST("/auth/login", h.LoginHabdler, middleware.GuestOnly)

	//protect
	e.GET("/test", h.TestHandler, middleware.AuthWiddleware)
}
