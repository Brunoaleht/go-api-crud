package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, productRoutes *ProductRoutes) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	productRoutes.InitRoutes(router)

}
