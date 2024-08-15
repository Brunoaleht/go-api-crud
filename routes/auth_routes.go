package routes

import (
	"go-api-commerce/controller"

	"github.com/gin-gonic/gin"
)

// AuthRoutes is a struct that represents auth routes
type AuthRoutes struct {
	Controller *controller.UserController
}

// NewAuthRoutes is a function to create a new Auth routes
func NewAuthRoutes(controller *controller.UserController) *AuthRoutes {
	return &AuthRoutes{
		Controller: controller,
	}
}

// InitRoutes is a function to initialize Auth routes
func (r *AuthRoutes) InitRoutes(router *gin.Engine) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/", r.Controller.Login)
		authRoutes.POST("/register", r.Controller.CreateUser)
		authRoutes.POST("/request-password", r.Controller.RequestUpdatePassword)
		authRoutes.PUT("/update-password/:id", r.Controller.UpdatePassword)
	}
}
