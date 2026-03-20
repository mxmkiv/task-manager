package main

import (
	"log"
	"task-manager/internal/config"
	"task-manager/internal/connection"
	"task-manager/internal/encoder"
	"task-manager/internal/handler"
	"task-manager/internal/repository"
	"task-manager/internal/routes"
	"task-manager/internal/service"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {

	config := config.NewServiceConfig()

	e := echo.New()
	// log all requests
	e.Use(middleware.RequestLogger())

	dbPool, dbErr := connection.DBConnection(config)
	if dbErr != nil {
		log.Fatal("database connection error", dbErr)
	}
	defer dbPool.Close()

	if PingErr := connection.PingDB(dbPool); PingErr != nil {
		log.Fatal("database ping error")
	}

	userEncoder := encoder.NewBcryptEncoder()
	repository := repository.NewUserRepository(dbPool)
	userService := service.NewUserService(repository, userEncoder)
	authService := service.NewAuthService(config.SecretKey, repository, userEncoder)
	handler := handler.NewHandler(authService, userService)

	routes.Setup(e, handler, config)

	err := e.Start(config.ServicePort)
	if err != nil {
		e.Logger.Error("server start error", "error", err)
	}
}
