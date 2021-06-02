package main

import (
	"github.com/IkezawaYuki/bookshelf-go/docs"
	"github.com/IkezawaYuki/bookshelf-go/src/infrastructure/http_handler"
	_ "github.com/swaggo/echo-swagger/example/docs"
	"os"
)

// @title bookshelf-api
// @version 0.0.0
// @description 読書●ーターを参考にした下位互換API
// @BasePath /v1
func main() {
	apiEndpoint := os.Getenv("API_ENDPOINT")
	if apiEndpoint != "" {
		docs.SwaggerInfo.Host = apiEndpoint
	} else {
		docs.SwaggerInfo.Host = "localhost:8080"
	}
	http_handler.StartApp()
}
