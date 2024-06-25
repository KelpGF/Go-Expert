package main

import (
	"fmt"

	"github.com/KelpGF/Go-Expert/08-APIs/configs"
)

func main() {
	configs := configs.LoadConfig(".")

	fmt.Println(configs.DBDriver)
	fmt.Println(configs.DBHost)
	fmt.Println(configs.DBPort)
	fmt.Println(configs.DBUser)
	fmt.Println(configs.DBPassword)
	fmt.Println(configs.DBName)
	fmt.Println(configs.WebServerPort)
	fmt.Println(configs.JWTSecret)
	fmt.Println(configs.JWTExpiresIn)
	fmt.Println(configs.TokenAuth)
}
