package middleware

import (
	"log"

	"github.com/labstack/echo/v5"
)

func TestMW(next echo.HandlerFunc) echo.HandlerFunc {

	return func(ctx *echo.Context) error {
		role := ctx.Request().Header.Get("user-role")
		if role == "admin" {
			log.Println("admin user detect")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil

	}

}
