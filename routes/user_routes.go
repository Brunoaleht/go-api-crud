package routes

import (
	"go-api-commerce/controller"

	"github.com/gin-gonic/gin"
)

// UserRoutes is a struct that represents user routes
type UserRoutes struct {
	Controller *controller.UserController
}

// NewUserRoutes is a function to create a new user routes
func NewUserRoutes(controller *controller.UserController) *UserRoutes {
	return &UserRoutes{
		Controller: controller,
	}
}

// InitRoutes is a function to initialize user routes
func (r *UserRoutes) InitRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		// Add other user routes here, like POST, PUT, DELETE, etc.
		userRoutes.DELETE("/:id", r.Controller.DeleteUser)
		userRoutes.PATCH("/:id", r.Controller.UpdateUser)
		userRoutes.GET("/:id", r.Controller.GetUserByID)
		userRoutes.POST("/", r.Controller.CreateUser)
		userRoutes.GET("/", r.Controller.GetUsers)
	}
}
