package infrastructure

import (
	"encoding/json"
	"github.com/IkezawaYuki/bookshelf-go/src/domain/model"
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
			token, err := getTokenFromRedis(key)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}
			if token == nil {
				return c.JSON(401, "login required")
			}
			return nil
		}
	}
}

func getTokenFromRedis(str string) (*model.Token, error) {
	jsonStr, err := RedisHandler.Get(str)
	if err != nil {
		return nil, err
	}
	var token model.Token
	if err := json.Unmarshal([]byte(jsonStr), &token); err != nil {
		return nil, err
	}
	return &token, nil
}
