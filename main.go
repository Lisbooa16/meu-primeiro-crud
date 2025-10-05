package main

import (
	"log"

	controller "github.com/Lisbooa16/meu-primeiro-crud/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("erro ao carregar .env file")
	}

	router := gin.Default()

	controller.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8000"); err != nil {
		log.Fatal(err)
	}

}
