package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func UserHandler(ctx *echo.Context) error {

	return ctx.String(http.StatusOK, "user data")

}

func TaskHandler(ctx *echo.Context) error {

	return ctx.String(http.StatusOK, "task data")

}

func MainHandler(ctx *echo.Context) error {
	return ctx.String(http.StatusOK, "Hello world")
}
