package main

import (
	"go-api-rest/controller"
	"go-api-rest/db"
	"go-api-rest/repository"
	"go-api-rest/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if(err != nil){
		panic(err)
	}
	
	//camada repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	//camada use case
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)

	//camada de controller
	productController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)
	server.Run(":8000")
}