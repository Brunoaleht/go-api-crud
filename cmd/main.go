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
	UserRepository := repository.NewUserRepository(dbConnection)
	AddressRepository := repository.NewAddressRepository(dbConnection)
	ProductRepository := repository.NewProductRepository(dbConnection)
	CategoryRepository := repository.NewCategoryRepository(dbConnection)
	CarRepository := repository.NewCarRepository(dbConnection)
	CarProductRepository := repository.NewCarProductRepository(dbConnection)

	//UseCase
	CategoryUseCase := usecase.NewCategoryUseCase(*CategoryRepository)
	ProductUseCase := usecase.NewProductUseCase(*ProductRepository, *CategoryRepository)
	UserUseCase := usecase.NewUserUseCase(*UserRepository)
	AddressUseCase := usecase.NewAddressUseCase(*AddressRepository)
	CarUseCase := usecase.NewCarUseCase(*CarRepository, *CarProductRepository)

	//Controller
	CategoryController := controller.NewCategoryController(*CategoryUseCase)
	ProductController := controller.NewProductController(*ProductUseCase)
	UserController := controller.NewUserController(*UserUseCase)
	AddressController := controller.NewAddressController(*AddressUseCase, *UserUseCase)
	CarController := controller.NewCarController(*CarUseCase)

	//Routes
	CarRoutes := routes.NewCarRoutes(CarController)
	AddressRoutes := routes.NewAddressRoutes(AddressController)
	CategoryRoutes := routes.NewCategoryRoutes(CategoryController)
	ProductRoutes := routes.NewProductRoutes(ProductController)
	UserRoutes := routes.NewUserRoutes(UserController)
	AuthRoutes := routes.NewAuthRoutes(UserController)

	//Init Routes
	routes.InitRoutes(server, ProductRoutes, UserRoutes, CategoryRoutes, AuthRoutes, AddressRoutes, CarRoutes)

	server.Run(":8000")
}
