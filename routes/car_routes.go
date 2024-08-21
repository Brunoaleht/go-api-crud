package routes

import (
	"go-api-commerce/controller"

	"github.com/gin-gonic/gin"
)

type CarRoutes struct {
	controller *controller.CarController
}

func NewCarRoutes(controller *controller.CarController) *CarRoutes {
	return &CarRoutes{
		controller: controller,
	}
}

func (r *CarRoutes) InitRoutes(router *gin.Engine) {
	carRoutes := router.Group("/cars")
	{
		// Get all cars or a specific car by ID
		carRoutes.GET("/", r.controller.GetCars)          // GET /cars - Get all cars
		carRoutes.GET("/:carId", r.controller.GetCarById) // GET /cars/:carId - Get car by ID

		// Get cars by user ID or get car with products by car ID
		carRoutes.GET("/user/:userId", r.controller.GetCarsByUserID)              // GET /cars/user/:userId - Get cars by user ID
		carRoutes.GET("/:carId/products", r.controller.GetCarWithProductsByCarID) // GET /cars/:carId/products - Get car with products by car ID

		// Add, update, and remove products from a car
		carRoutes.POST("/user/:userId/products", r.controller.AddProductToCar)                              // POST /cars/user/:userId/products - Add product to car
		carRoutes.PUT("/user/:userId/products", r.controller.UpdateProductQuantity)                         // PUT /cars/user/:userId/products - Update product quantity in car
		carRoutes.DELETE("/:cartId/user/:userId/products/:carProductId", r.controller.RemoveProductFromCar) // DELETE /cars/user/:userId/:cartId/products/:carProductId - Remove product from car
	}
}
