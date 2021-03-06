package auth

import (
	"encoding/json"
	"github.com/IkezawaYuki/bookshelf-go/src/domain/model"
	"github.com/IkezawaYuki/bookshelf-go/src/infrastructure/redis"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func Guard() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Authorizationヘッダーからキーを取得
			auth := c.Request().Header.Get("Authorization")
			if auth == "" {
				return c.JSON(401, "login required")
			}
			key := strings.ReplaceAll(auth, "Bearer ", "")

			// redisからログイン情報を取得し、contextにセット
			token, err := getTokenFromRedis(key)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}
			if token == nil && token.RefreshToken == "" {
				return c.JSON(401, "login required")
			}
			c.Set("refresh_token", token.RefreshToken)
			return next(c)
		}
	}
}

func getTokenFromRedis(str string) (*model.Token, error) {
	jsonStr, err := redis.Handler.Get(str)
	if err != nil {
		return nil, err
	}
	var token model.Token
	if err := json.Unmarshal([]byte(jsonStr), &token); err != nil {
		return nil, err
	}
	return &token, nil
}
