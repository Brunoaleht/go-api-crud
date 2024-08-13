package controller

import (
	"go-api-commerce/model"
	"go-api-commerce/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProductController is a struct that represents a product controller
type ProductController struct {
	ProductUseCase usecase.ProductUseCase
}

// NewProductController is a function to create a new product controller
func NewProductController(usecase usecase.ProductUseCase) *ProductController {
	return &ProductController{
		ProductUseCase: usecase,
	}
}

// GetProducts is a function to get all products
func (pc *ProductController) GetProducts(ctx *gin.Context) {
	response := pc.ProductUseCase.GetProducts()
	if !response.Success {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":  response.Message,
			"success":  response.Success,
			"products": response.Data,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"products": response.Data,
		"message":  response.Message,
		"success":  response.Success,
	})
}

// CreateProduct is a function to create a new product
func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"success": false,
		})
		return
	}

	response := pc.ProductUseCase.CreateProduct(product)
	if !response.Success {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":  response.Message,
			"success":  response.Success,
			"products": response.Data,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"product": response.Data,
		"message": response.Message,
		"success": response.Success,
	})
}

// GetProductByID is a function to get a product by ID
func (pc *ProductController) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
			"success": false,
		})
		return
	}
	response := pc.ProductUseCase.GetProductByID(idNumber)
	if !response.Success {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": response.Message,
			"success": response.Success,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"product": response.Data,
		"message": response.Message,
		"success": response.Success,
	})
}

// UpdateProduct is a function to update a product
func (pc *ProductController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
			"success": false,
		})
		return
	}
	var product model.Product

	product.ID = idNumber

	err = ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"success": false,
		})
		return
	}

	response := pc.ProductUseCase.UpdateProduct(product)
	if !response.Success {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":  response.Message,
			"success":  response.Success,
			"products": response.Data,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"product": response.Data,
		"message": response.Message,
		"success": response.Success,
	})

}

// DeleteProduct is a function to delete a product
func (pc *ProductController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
			"success": false,
		})
		return
	}

	response := pc.ProductUseCase.DeleteProduct(idNumber)
	if !response.Success {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":  response.Message,
			"success":  response.Success,
			"products": response.Data,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"product": response.Data,
		"message": response.Message,
		"success": response.Success,
	})
}
