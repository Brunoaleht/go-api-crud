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
	UserRepository := repository.NewUserRepository(dbConnection)
	CategoryRepository := repository.NewCategoryRepository(dbConnection)

	//UseCase
	CategoryUseCase := usecase.NewCategoryUseCase(*CategoryRepository)
	ProductUseCase := usecase.NewProductUseCase(*ProductRepository, *CategoryRepository)
	UserUseCase := usecase.NewUserUseCase(*UserRepository)

	//Controller
	CategoryController := controller.NewCategoryController(*CategoryUseCase)
	ProductController := controller.NewProductController(*ProductUseCase)
	UserController := controller.NewUserController(*UserUseCase)

	//Routes
	CategoryRoutes := routes.NewCategoryRoutes(CategoryController)
	ProductRoutes := routes.NewProductRoutes(ProductController)
	UserRoutes := routes.NewUserRoutes(UserController)

	//Init Routes
	routes.InitRoutes(server, ProductRoutes, UserRoutes, CategoryRoutes)

	server.Run(":8000")
}
