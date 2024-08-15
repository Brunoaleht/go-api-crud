package controller

import (
	"go-api-commerce/model"
	"go-api-commerce/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CategoryController is a struct that represents a category controller
type CategoryController struct {
	CategoryUseCase usecase.CategoryUseCase
}

// NewCategoryController is a function to create a new category controller
func NewCategoryController(usecase usecase.CategoryUseCase) *CategoryController {
	return &CategoryController{
		CategoryUseCase: usecase,
	}
}

// GetCategories is a function to get all categories
func (cc *CategoryController) GetCategories(ctx *gin.Context) {
	response := cc.CategoryUseCase.GetCategories()
	if !response.Success {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":    response.Message,
			"success":    response.Success,
			"categories": response.Data,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"categories": response.Data,
		"message":    response.Message,
		"success":    response.Success,
	})
}

// CreateCategory is a function to create a new category
func (cc *CategoryController) CreateCategory(ctx *gin.Context) {
	var category model.Category
	err := ctx.BindJSON(&category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"success": false,
		})
		return
	}

	response := cc.CategoryUseCase.CreateCategory(category)
	if !response.Success {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": response.Message,
			"success": response.Success,
		})
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message":  response.Message,
		"success":  response.Success,
		"category": response.Data,
	})
}

// GetCategoryByID is a function to get a category by ID
func (cc *CategoryController) GetCategoryByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"success": false,
		})
		return
	}

	response := cc.CategoryUseCase.GetCategoryByID(idNumber)
	if !response.Success {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": response.Message,
			"success": response.Success,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"category": response.Data,
		"message":  response.Message,
		"success":  response.Success,
	})
}

// UpdateCategory is a function to update a category
func (cc *CategoryController) UpdateCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"success": false,
		})
		return
	}

	var category model.Category

	category.ID = idNumber

	err = ctx.BindJSON(&category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"success": false,
		})
		return
	}

	response := cc.CategoryUseCase.UpdateCategory(category)
	if !response.Success {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": response.Message,
			"success": response.Success,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"category": response.Data,
		"message":  response.Message,
		"success":  response.Success,
	})

}

// DeleteCategory is a function to delete a category
func (cc *CategoryController) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"success": false,
		})
		return
	}

	response := cc.CategoryUseCase.DeleteCategory(idNumber)
	if !response.Success {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": response.Message,
			"success": response.Success,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  response.Message,
		"success":  response.Success,
		"category": response.Data,
	})
}
