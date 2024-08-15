package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, productRoutes *ProductRoutes, userRoutes *UserRoutes, categoryRoutes *CategoryRoutes, authRoutes *AuthRoutes) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	productRoutes.InitRoutes(router)
	userRoutes.InitRoutes(router)
	categoryRoutes.InitRoutes(router)
	authRoutes.InitRoutes(router)
}
