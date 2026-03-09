package middleware

import (
	"net/http"
	"strings"
	"task-manager/internal/auth"

	"github.com/labstack/echo/v5"
)

func AuthWiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c *echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		// bearer check
		if !strings.HasPrefix(authHeader, "Bearer") {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization header format, expected 'Bearer <token>'")
		}

		token := authHeader[7:]
		if err := auth.ParseToken(token); err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}

		return next(c)
	}
}

func GuestOnly(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c *echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return next(c)
		}

		token := authHeader[7:]
		if err := auth.ParseToken(token); err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "access denied: user already authenticated")
		}

		return next(c)
	}

}
