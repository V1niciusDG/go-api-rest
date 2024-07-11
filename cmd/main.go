package main

import (
	"go-api-rest/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	//camada de controller
	productController := controller.NewProductController()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)
	server.Run(":8000")
}