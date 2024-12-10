package main

import (
	"github.com/gin-gonic/gin"
	"practise.com/rest-api-go/db"
	"practise.com/rest-api-go/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost
}
