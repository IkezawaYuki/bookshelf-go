package infrastructure

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func AuthGuard() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")
			if auth == "" {
				return c.JSON(401, "login required")
			}
			key := strings.ReplaceAll(auth, "Bearer ", "")
			token, err := RedisHandler.Get(key)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}
			if token == "" {
				return c.JSON(401, "login required")
			}

			return nil
		}
	}
}
