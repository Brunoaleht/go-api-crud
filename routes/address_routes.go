package routes

import (
	"go-api-commerce/controller"

	"github.com/gin-gonic/gin"
)

// AddressRoutes is a struct that represents address routes
type AddressRoutes struct {
	controller *controller.AddressController
}

// NewAddressRoutes is a function to create a new address routes
func NewAddressRoutes(controller *controller.AddressController) *AddressRoutes {
	return &AddressRoutes{
		controller: controller,
	}
}

// InitRoutes is a function to initialize address routes
func (r *AddressRoutes) InitRoutes(router *gin.Engine) {
	addressRouter := router.Group("/addresses")
	{
		addressRouter.GET("/users/:userId", r.controller.GetAddressByUserID)
		addressRouter.DELETE("/:id", r.controller.DeleteAddress)
		addressRouter.PATCH("/:id", r.controller.UpdateAddress)
		addressRouter.GET("/:id", r.controller.GetAddressByID)
		addressRouter.POST("/", r.controller.CreateAddress)
	}
}
