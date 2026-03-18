package routes

import (
	"task-manager/internal/config"
	"task-manager/internal/handler"
	"task-manager/internal/middleware"

	"github.com/labstack/echo/v5"
)

func Setup(e *echo.Echo, h *handler.Handler, config *config.ServiceConfig) {

	//api := e.Group("/api/v1")

	//auth
	e.POST("/auth/register", h.RegisterHandler, middleware.GuestOnly(config.SecretKey))
	e.POST("/auth/login", h.LoginHandler, middleware.GuestOnly(config.SecretKey))
	e.POST("/admin", h.CreateAdmin, middleware.GuestOnly(config.SecretKey))

	//protect
	protect := e.Group("", middleware.AuthWiddleware(config.SecretKey))
	//users
	protect.GET("/users", h.AllUsersHandler, middleware.AdminOnly)
	protect.GET("/users/:id", h.UserById, middleware.AdminOnly)
	protect.PATCH("/users/:id", h.UpdateUserDataHandler, middleware.PermissionMiddleware(config.SecretKey))
	protect.DELETE("/users/:id", h.DeleteUserHandler, middleware.PermissionMiddleware(config.SecretKey))
	//tasks
}
