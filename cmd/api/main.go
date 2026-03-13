package main

import (
	"context"
	"log"
	"task-manager/internal/config"
	"task-manager/internal/encoder"
	"task-manager/internal/handler"
	"task-manager/internal/repository"
	"task-manager/internal/routes"
	"task-manager/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {

	config := config.NewServiceConfig()

	e := echo.New()
	// log all requests
	e.Use(middleware.RequestLogger())

	dbPool, dbErr := pgxpool.New(context.Background(), config.DbConnectionUrl)
	if dbErr != nil {
		log.Fatal("database connection error", dbErr)
	}

	userEncoder := encoder.NewBcryptEncoder()
	repository := repository.NewUserRepository(dbPool)
	userService := service.NewUserService(repository, userEncoder)
	authService := service.NewAuthService(config.SecretKey, repository, userEncoder)
	handler := handler.NewHandler(authService, userService)

	// services
	// logger

	routes.Setup(e, handler, config)

	err := e.Start(config.ServicePort)
	if err != nil {
		e.Logger.Error("server start error", "error", err)
	}
}
