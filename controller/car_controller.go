package controller

import (
	"go-api-commerce/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CarController struct {
	cu usecase.CarUseCase
}

func NewCarController(cu usecase.CarUseCase) *CarController {
	return &CarController{
		cu: cu,
	}
}

func (cc *CarController) GetCars(ctx *gin.Context) {
	response := cc.cu.GetCars()
	if !response.Success {
		ctx.JSON(http.StatusFound, gin.H{
			"message": response.Message,
			"success": response.Success,
			"cars":    response.Data,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"cars":    response.Data,
		"message": response.Message,
		"success": response.Success,
	})
}

func (cc *CarController) GetCarById(ctx *gin.Context) {
	carID, err := strconv.Atoi(ctx.Param("carId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid car ID",
			"success": false,
		})
		return
	}
	response := cc.cu.GetCarByID(carID)
	if !response.Success {
		ctx.JSON(http.StatusFound, gin.H{
			"message": response.Message,
			"success": response.Success,
			"car":     response.Data,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car":     response.Data,
		"message": response.Message,
		"success": response.Success,
	})
}

func (cc *CarController) GetCarsByUserID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID",
			"success": false,
		})
		return
	}
	response := cc.cu.GetCarsByUserID(userID)
	if !response.Success {
		ctx.JSON(http.StatusFound, gin.H{
			"message": response.Message,
			"success": response.Success,
			"cars":    response.Data,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"cars":    response.Data,
		"message": response.Message,
		"success": response.Success,
	})
}

func (cc *CarController) GetCarWithProductsByCarID(ctx *gin.Context) {
	carID, err := strconv.Atoi(ctx.Param("carId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid car ID",
			"success": false,
		})
		return
	}
	response := cc.cu.GetCarWithProductsByCarID(carID)
	if !response.Success {
		ctx.JSON(http.StatusFound, gin.H{
			"message": response.Message,
			"success": response.Success,
			"car":     response.Data,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car":     response.Data,
		"message": response.Message,
		"success": response.Success,
	})
}

func (cc *CarController) AddProductToCar(ctx *gin.Context) {

}
