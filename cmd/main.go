package main

import (
	"go-api-commerce/controller"
	"go-api-commerce/db"
	"go-api-commerce/repository"
	"go-api-commerce/routes"
	"go-api-commerce/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	//UseCase
	ProductUseCase := usecase.NewProductUseCase(*ProductRepository)

	//Controller
	ProductController := controller.NewProductController(*ProductUseCase)

	//Routes
	ProductRoutes := routes.NewProductRoutes(ProductController)

	//Init Routes
	routes.InitRoutes(server, ProductRoutes)

	server.Run(":8000")
}
