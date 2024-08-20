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
	pu usecase.ProductUseCase
}

func NewCarController(cu usecase.CarUseCase, pu usecase.ProductUseCase) *CarController {
	return &CarController{
		cu: cu,
		pu: pu,
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

	// Verifica se o produto existe e se tem estoque
	responseProduct := cc.pu.GetProductByID(carProduct.ProductID)
	if !responseProduct.Success {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": responseProduct.Message,
			"success": responseProduct.Success,
		})
		return
	}

	product := responseProduct.Data
	if product.StockQuantity == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Product out of stock",
			"success": false,
		})
		return
	}

	carProduct.UnitPrice = product.Price

	// Atualiza a quantidade de estoque
	product.StockQuantity -= carProduct.Quantity
	responseUpdateProduct := cc.pu.UpdateProduct(product)
	if !responseUpdateProduct.Success {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": responseUpdateProduct.Message,
			"success": responseUpdateProduct.Success,
		})
		return
	}

	// Adiciona o produto ao carrinho, cria o carrinho se necessário
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
