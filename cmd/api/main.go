package main

import (
	"context"
	"log"
	"task-manager/internal/config"
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

	userRepo := repository.NewUserRepository(dbPool)
	srv := service.NewAuthService(config.SecretKey, userRepo)
	handler := handler.NewHandler(srv)

	// services
	// logger

	routes.Setup(e, handler)

	err := e.Start(config.ServicePort)
	if err != nil {
		e.Logger.Error("server start error", "error", err)
	}
}
