package http_handler

import (
	"github.com/IkezawaYuki/bookshelf-go/src/infrastructure/auth"
	"github.com/IkezawaYuki/bookshelf-go/src/infrastructure/redis"
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/controller"
	"github.com/IkezawaYuki/bookshelf-go/src/logger"
	"github.com/IkezawaYuki/bookshelf-go/src/registry"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/swaggo/echo-swagger"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

func StartApp() {
	gomniauth.SetSecurityKey(os.Getenv("SECURITY_KEY"))
	gomniauth.WithProviders(
		google.New(
			os.Getenv("CLIENT_ID"),
			os.Getenv("CLIENT_SECRET"),
			os.Getenv("REDIRECT_URL"),
		),
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

	var allowsOrigins = []string{
		"http://localhost:3000",
	}

	e.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: allowsOrigins,
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodDelete,
				http.MethodPatch,
			},
			AllowHeaders: []string{
				echo.HeaderAccessControlAllowHeaders,
				echo.HeaderContentType,
				echo.HeaderContentLength,
				echo.HeaderAcceptEncoding,
				echo.HeaderXCSRFToken,
				echo.HeaderAuthorization,
			},
			AllowCredentials: true,
			MaxAge:           86400,
		}))

	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	/*
		認証の必要がないAPI
	*/
	e.GET("/v1/version", func(c echo.Context) error {
		return bookShelfCtr.GetVersion(c)
	})

	/*
		認証API
	*/
	e.GET("/v1/auth/login", func(c echo.Context) error {
		return authCtr.Login(c)
	})

	e.GET("/v1/auth/callback", func(c echo.Context) error {
		return authCtr.Callback(c)
	})

	e.GET("/v1/auth/logout", func(c echo.Context) error {
		key := c.Request().Header.Get(echo.HeaderAuthorization)
		key = strings.ReplaceAll(key, "Bearer ", "")
		if key == "" {
			return nil
		}
		c.Set("key", key)
		return authCtr.Logout(c)
	})

	/*
		認証が必要なAPI
	*/
	g := e.Group("/v1")
	g.Use(auth.Guard())

	g.GET("/book/:id", func(c echo.Context) error {
		return bookShelfCtr.GetBook(c)
	})

	g.GET("/books", func(c echo.Context) error {
		return bookShelfCtr.GetBooks(c)
	})

	g.POST("/book", func(c echo.Context) error {
		return bookShelfCtr.CreateBook(c)
	})

	g.PATCH("/book", func(c echo.Context) error {
		return bookShelfCtr.UpdateBook(c)
	})

	g.GET("/book/detail/:id", func(c echo.Context) error {
		return bookShelfCtr.ShowBook(c)
	})

	g.DELETE("/book/:id", func(c echo.Context) error {
		return bookShelfCtr.DeleteBook(c)
	})

	g.GET("/user/detail/:id", func(c echo.Context) error {
		return bookShelfCtr.ShowUser(c)
	})

	g.GET("/users", func(c echo.Context) error {
		return bookShelfCtr.GetUsers(c)
	})

	g.GET("/users/report", func(c echo.Context) error {
		return bookShelfCtr.OutputUsersReport(c)
	})

	g.GET("/review/:id", func(c echo.Context) error {
		return bookShelfCtr.GetReview(c)
	})

	g.POST("/review", func(c echo.Context) error {
		return bookShelfCtr.CreateReview(c)
	})

	g.PATCH("/review", func(c echo.Context) error {
		return bookShelfCtr.UpdateReview(c)
	})

	g.DELETE("/review/:id", func(c echo.Context) error {
		return bookShelfCtr.DeleteReview(c)
	})

	g.GET("/comment/:id", func(c echo.Context) error {
		return bookShelfCtr.GetComment(c)
	})

	g.POST("/comment", func(c echo.Context) error {
		return bookShelfCtr.CreateComment(c)
	})

	g.PATCH("/comment", func(c echo.Context) error {
		return bookShelfCtr.UpdateComment(c)
	})

	g.DELETE("/comment/:id", func(c echo.Context) error {
		return bookShelfCtr.DeleteComment(c)
	})

	go func() {
		if err := e.Start(":8080"); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("Shutdown...")
	_ = container.Clean()
	_ = redis.Handler.Close()
}
