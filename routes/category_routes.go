package routes

import (
	"go-api-commerce/controller"

	"github.com/gin-gonic/gin"
)

// InitRoutes is a function to initialize all routes
type CategoryRoutes struct {
	Controller *controller.CategoryController
}

func NewCategoryRoutes(controller *controller.CategoryController) *CategoryRoutes {
	return &CategoryRoutes{
		Controller: controller,
	}
}

func (r *CategoryRoutes) InitRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")
	{
		// Add other category routes here, like POST, PUT, DELETE, etc.
		categoryRoutes.DELETE("/:id", r.Controller.DeleteCategory)
		categoryRoutes.PATCH("/:id", r.Controller.UpdateCategory)
		categoryRoutes.GET("/:id", r.Controller.GetCategoryByID)
		categoryRoutes.POST("/", r.Controller.CreateCategory)
		categoryRoutes.GET("/", r.Controller.GetCategories)
	}
}
