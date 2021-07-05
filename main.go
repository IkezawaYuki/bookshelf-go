package main

import "fmt"

//
//// @title bookshelf-api
//// @version 0.0.0
//// @description 読書●ーターを参考にした下位互換API
//// @BasePath /v1
//func main() {
//	docs.SwaggerInfo.Host = os.Getenv("API_ENDPOINT")
//	if docs.SwaggerInfo.Host != "" {
//		docs.SwaggerInfo.Host = "localhost:8080"
//	}
//	http_handler.StartApp()
//}

func main() {

	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Print("Fizz")
		} else if i%5 == 0 {
			fmt.Print("Buzz")
		}
		if i%7 == 0 {
			fmt.Print("Hoo")
		}
		if i%11 == 0 {
			fmt.Print("Bar")
		}
		if i%3 != 0 && i%5 != 0 && i%7 != 0 && i%11 != 0 {
			fmt.Print(i)
		}
		fmt.Println()
	}
}
