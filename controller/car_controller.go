package controller

import (
	"go-api-commerce/model"
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
	userID, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ID do usuário inválido",
			"success": false,
		})
		return
	}

	var carProduct model.CarProduct
	err = ctx.BindJSON(&carProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Dados do produto no carro inválidos",
			"success": false,
		})
		return
	}

	// Verificar a existência do produto e estoque na camada de use case
	response := cc.cu.AddProductToCar(userID, carProduct)
	if !response.Success {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": response.Message,
			"success": response.Success,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": response.Message,
		"success": response.Success,
		"car":     response.Data,
	})
}

func (cc *CarController) RemoveProductFromCar(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID",
			"success": false,
		})
		return
	}

	carID, err := strconv.Atoi(ctx.Param("cartId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid car ID",
			"success": false,
		})
		return
	}

	carProductID, err := strconv.Atoi(ctx.Param("carProductId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid car product ID",
			"success": false,
		})
		return
	}

	response := cc.cu.RemoveProductFromCar(userID, carID, carProductID)
	if !response.Success {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": response.Message,
			"success": response.Success,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": response.Message,
		"success": response.Success,
	})
}

func (cc *CarController) UpdateProductQuantity(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID",
			"success": false,
		})
		return
	}

	var carProduct model.CarProduct
	err = ctx.BindJSON(&carProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid car product data",
			"success": false,
		})
		return
	}

	response := cc.cu.UpdateProductQuantity(userID, carProduct)
	if !response.Success {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": response.Message,
			"success": response.Success,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": response.Message,
		"success": response.Success,
	})
}
