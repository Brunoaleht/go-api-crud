package routes

import (
	"go-api-commerce/controller"

	"github.com/gin-gonic/gin"
)

type ProductRoutes struct {
	Controller *controller.ProductController
}

func NewProductRoutes(controller *controller.ProductController) *ProductRoutes {
	return &ProductRoutes{
		Controller: controller,
	}
}

func (r *ProductRoutes) InitRoutes(router *gin.Engine) {
	productRoutes := router.Group("/products")
	{
		// Adicione outras rotas de produtos aqui, como POST, PUT, DELETE, etc.
		productRoutes.DELETE("/:id", r.Controller.DeleteProduct)
		productRoutes.PATCH("/:id", r.Controller.UpdateProduct)
		productRoutes.GET("/category/:id", r.Controller.GetProductsByCategoryID)
		productRoutes.GET("/:id", r.Controller.GetProductByID)
		productRoutes.POST("/", r.Controller.CreateProduct)
		productRoutes.GET("/", r.Controller.GetProducts)
	}
}
