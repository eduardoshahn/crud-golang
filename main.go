package main

import (
	//"fmt"
	"log"
	//"os"

	"github.com/eduardoshahn/crud-golang.git/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	// Se ja tiver uma aplicação na 8080 irá dar erro
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}