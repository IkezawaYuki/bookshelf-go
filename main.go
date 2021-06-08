package main

import (
	"github.com/IkezawaYuki/bookshelf-go/docs"
	"github.com/IkezawaYuki/bookshelf-go/src/infrastructure/http_handler"
	"os"
)

// @title bookshelf-api
// @version 0.0.0
// @description 読書●ーターを参考にした下位互換API
// @BasePath /v1
func main() {
	docs.SwaggerInfo.Host = os.Getenv("API_ENDPOINT")
	if docs.SwaggerInfo.Host != "" {
		docs.SwaggerInfo.Host = "localhost:8080"
	}
	http_handler.StartApp()
}
