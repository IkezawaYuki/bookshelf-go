package main

import (
	"github.com/IkezawaYuki/bookshelf-go/src/infrastructure"
	"os"
)

// @title bookshelf-api
// @version 0.0.0
// @description aaa
// @BasePath /v1
func main() {
	apiEndpoint := os.Getenv("API_ENDPOINT")
	if apiEndpoint != "" {

	} else {

	}
	infrastructure.StartApp()
}
