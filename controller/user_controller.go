package controller

import (
	"go-api-commerce/model"
	"go-api-commerce/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserController is a struct that represents a user controller
type UserController struct {
	UserUseCase usecase.UserUseCase
}

func NewUserController(usecase usecase.UserUseCase) *UserController {
	return &UserController{
		UserUseCase: usecase,
	}
}

// GetUsers is a function to get all users
func (uc *UserController) GetUsers(ctx *gin.Context) {
	response := uc.UserUseCase.GetUsers()
	if !response.Success {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": response.Message,
			"success": response.Success,
			"users":   response.Data,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users":   response.Data,
		"message": response.Message,
		"success": response.Success,
	})
}

// CreateUser is a function to create a new user
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"success": false,
		})
		return
	}

	response := uc.UserUseCase.CreateUser(user)
	if !response.Success {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": response.Message,
			"success": response.Success,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"user":    response.Data,
		"message": response.Message,
		"success": response.Success,
	})
}

// GetUserByID is a function to get a user by ID
func (uc *UserController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
			"success": false,
		})
		return
	}

	response := uc.UserUseCase.GetUserByID(idNumber)
	if !response.Success {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": response.Message,
			"success": response.Success,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user":    response.Data,
		"message": response.Message,
		"success": response.Success,
	})
}

// UpdateUser is a function to update a user
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
			"success": false,
		})
		return
	}

	var user model.User
	user.ID = idNumber

	err = ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"success": false,
		})
		return
	}

	response := uc.UserUseCase.UpdateUser(user)
	if !response.Success {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": response.Message,
			"success": response.Success,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user":    response.Data,
		"message": response.Message,
		"success": response.Success,
	})
}

// DeleteUser is a function to delete a user
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
			"success": false,
		})
		return
	}

	response := uc.UserUseCase.DeleteUser(idNumber)
	if !response.Success {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": response.Message,
			"success": response.Success,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user":    response.Data,
		"message": response.Message,
		"success": response.Success,
	})
}
