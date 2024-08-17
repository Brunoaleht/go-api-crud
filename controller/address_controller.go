package controller

import (
	"go-api-commerce/model"
	"go-api-commerce/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AddressController struct {
	usecaseAddress usecase.AddressUseCase
	usecaseUser    usecase.UserUseCase
}

func NewAddressController(usecaseAddress usecase.AddressUseCase, usecaseUser usecase.UserUseCase) *AddressController {
	return &AddressController{
		usecaseAddress: usecaseAddress,
		usecaseUser:    usecaseUser,
	}
}

func (ac *AddressController) GetAddressByUserID(ctx *gin.Context) {
	userID := ctx.Param("userId")
	userIDNumber, err := strconv.Atoi(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user id",
			"success": false,
		})
		return
	}

	responseUser := ac.usecaseUser.GetUserByID(userIDNumber)
	if !responseUser.Success {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
			"success": false,
		})
		return
	}

	response := ac.usecaseAddress.GetAddressByUserID(userIDNumber)
	if !response.Success {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message":   response.Message,
			"addresses": response.Data,
			"success":   response.Success,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"addresses": response.Data,
		"message":   response.Message,
		"success":   response.Success,
	})
}

func (ac *AddressController) CreateAddress(ctx *gin.Context) {
	var address model.Address
	err := ctx.BindJSON(&address)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"success": false,
		})
		return
	}

	response := ac.usecaseAddress.CreateAddress(address)
	if !response.Success {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": response.Message,
			"success": response.Success,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": response.Message,
		"success": response.Success,
		"address": response.Data,
	})
}

func (ac *AddressController) GetAddressByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"success": false,
		})
		return
	}

	response := ac.usecaseAddress.GetAddressByID(idNumber)
	if !response.Success {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": response.Message,
			"success": response.Success,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"address": response.Data,
		"message": response.Message,
		"success": response.Success,
	})
}

func (ac *AddressController) UpdateAddress(ctx *gin.Context) {
	id := ctx.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"success": false,
		})
		return
	}

	var address model.Address

	address.ID = idNumber

	err = ctx.BindJSON(&address)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"success": false,
		})
		return
	}

	response := ac.usecaseAddress.UpdateAddress(address)
	if !response.Success {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": response.Message,
			"success": response.Success,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"address": response.Data,
		"message": response.Message,
		"success": response.Success,
	})
}

func (ac *AddressController) DeleteAddress(ctx *gin.Context) {
	id := ctx.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"success": false,
		})
		return
	}

	responseAddress := ac.usecaseAddress.GetAddressByID(idNumber)
	if !responseAddress.Success {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": responseAddress.Message,
			"success": responseAddress.Success,
		})
		return
	}

	response := ac.usecaseAddress.DeleteAddress(idNumber)
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
		"address": responseAddress.Data,
	})
}
