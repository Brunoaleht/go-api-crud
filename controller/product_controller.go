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
	//example:
	products, err := pc.ProductUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error" + err.Error(),
			"success": false,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"products": products,
		"message":  "success",
		"success":  true,
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

	createdProduct, err := pc.ProductUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error" + err.Error(),
			"success": false,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"product": createdProduct,
		"message": "Created successfully",
		"success": true,
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
	product, err := pc.ProductUseCase.GetProductByID(idNumber)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
			"success": false,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"product": product,
		"message": "success",
		"success": true,
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

	_, err = pc.ProductUseCase.GetProductByID(idNumber)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
			"success": false,
		})
		return
	}

	if product.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Name is required",
			"success": false,
		})
		return
	}

	if product.Price == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Price is required",
			"success": false,
		})
		return
	}

	updatedProduct, err := pc.ProductUseCase.UpdateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error" + err.Error(),
			"success": false,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"product": updatedProduct,
		"message": "Updated successfully",
		"success": true,
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

	product, err := pc.ProductUseCase.GetProductByID(idNumber)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
			"success": false,
		})
		return
	}

	err = pc.ProductUseCase.DeleteProduct(idNumber)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error" + err.Error(),
			"success": false,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"product": product,
		"message": "Deleted successfully",
		"success": true,
	})
}
