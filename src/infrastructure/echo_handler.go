package infrastructure

import (
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/controller"
	"github.com/IkezawaYuki/bookshelf-go/src/registry"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"os"
	"os/signal"
	"strings"
)

func StartApp() {
	gomniauth.SetSecurityKey(os.Getenv("SECURITY_KEY"))
	gomniauth.WithProviders(
		google.New(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("REDIRECT_URL")),
	)
	container, err := registry.NewContainer()
	if err != nil {
		panic(err)
	}
	bookShelfCtr := container.Resolve("bookshelf-controller").(*controller.BookshelfController)
	authCtr := container.Resolve("auth-controller").(*controller.AuthController)

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Request().URL.String(), "v1/version") {
				return true
			}
			return false
		},
	}))

	/*
		認証の必要がないAPI
	*/
	e.GET("/v1/version", func(c echo.Context) error {
		return bookShelfCtr.GetVersion(c)
	})

	/*
		認証API
	*/
	e.GET("v1/auth/login", func(c echo.Context) error {
		return authCtr.Login(c)
	})

	e.GET("v1/auth/callback", func(c echo.Context) error {
		return authCtr.Callback(c)
	})

	e.GET("v1/auth/logout", func(c echo.Context) error {
		key := c.Request().Header.Get("Authentication")
		if key == "" {
			return nil
		}
		c.Set("key", key)
		return authCtr.Logout(c)
	})

	/*
		認証が必要なAPI
	*/

	go func() {
		if err := e.Start(":8080"); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	_ = container.Clean()
	_ = RedisHandler.Close()
	// todo DB close
}
