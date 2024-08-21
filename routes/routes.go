package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, productRoutes *ProductRoutes, userRoutes *UserRoutes, categoryRoutes *CategoryRoutes, authRoutes *AuthRoutes, addressRoutes *AddressRoutes, carRoutes *CarRoutes) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//
	carRoutes.InitRoutes(router)

	//
	categoryRoutes.InitRoutes(router)
	productRoutes.InitRoutes(router)

	//
	addressRoutes.InitRoutes(router)
	userRoutes.InitRoutes(router)
	authRoutes.InitRoutes(router)
}
