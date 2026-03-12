package middleware

import (
	"net/http"
	"strings"
	"task-manager/internal/auth"

	"github.com/labstack/echo/v5"
)

func AuthWiddleware(secretKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			// bearer check
			if !strings.HasPrefix(authHeader, "Bearer ") {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization header format, expected 'Bearer <token>'")
			}

			token := authHeader[7:]
			claims, err := auth.ParseToken(token, secretKey)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}

			c.Set("userRole", claims.Role)

			return next(c)

		}
	}
}

func GuestOnly(secretKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				return next(c)
			}

			token := authHeader[7:]
			if _, err := auth.ParseToken(token, secretKey); err == nil {
				return echo.NewHTTPError(http.StatusForbidden, "access denied: user already authenticated")
			}

			return next(c)
		}
	}
}

func AdminOnly(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c *echo.Context) error {
		role := c.Get("userRole")
		if role != "admin" {
			return echo.NewHTTPError(http.StatusForbidden, "access denied")
		}
		return next(c)
	}

}
