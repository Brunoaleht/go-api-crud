package usecase

import (
	"go-api-commerce/model"
	"go-api-commerce/repository"
	"go-api-commerce/utils"
)

type IResponseListAddress struct {
	Message string          `json:"message"`
	Data    []model.Address `json:"data"`
	Success bool            `json:"success"`
}

type IResponseAddress struct {
	Message string        `json:"message"`
	Data    model.Address `json:"data"`
	Success bool          `json:"success"`
}

type AddressUseCase struct {
	repo repository.AddressRepository
}

func NewAddressUseCase(repo repository.AddressRepository) *AddressUseCase {
	return &AddressUseCase{
		repo: repo,
	}
}

func (uc *AddressUseCase) GetAddressByUserID(userID int) IResponseListAddress {
	address, err := uc.repo.GetAddressByUserID(userID)
	if err != nil {
		return IResponseListAddress{
			Message: "Error getting address" + err.Error(),
			Data:    []model.Address{},
			Success: false,
		}
	}

	return IResponseListAddress{
		Message: "Success getting address",
		Data:    address,
		Success: true,
	}

}

func (uc *AddressUseCase) CreateAddress(address model.Address) IResponseAddress {
	id, err := uc.repo.CreateAddress(address)
	if err != nil {
		return IResponseAddress{
			Message: "Error creating address" + err.Error(),
			Data:    model.Address{},
			Success: false,
		}
	}

	address.ID = id
	return IResponseAddress{
		Message: "Success creating address",
		Data:    address,
		Success: true,
	}
}

func (uc *AddressUseCase) GetAddressByID(id int) IResponseAddress {
	address, err := uc.repo.GetAddressByID(id)
	if err != nil {
		return IResponseAddress{
			Message: "Error getting address" + err.Error(),
			Data:    model.Address{},
			Success: false,
		}
	}

	return IResponseAddress{
		Message: "Success getting address",
		Data:    address,
		Success: true,
	}
}

func (uc *AddressUseCase) UpdateAddress(address model.Address) IResponseAddress {
	foundAddress, err := uc.repo.GetAddressByID(address.ID)
	if err != nil {
		return IResponseAddress{
			Message: "Error getting address" + err.Error(),
			Data:    model.Address{},
			Success: false,
		}
	}

	if utils.IntToBool(foundAddress.UserID) {
		return IResponseAddress{
			Message: "Address to user not found",
			Data:    model.Address{},
			Success: false,
		}
	}

	if address.UserID == 0 {
		address.UserID = foundAddress.UserID
	}

	err = uc.repo.UpdateAddress(address)
	if err != nil {
		return IResponseAddress{
			Message: "Error updating address" + err.Error(),
			Data:    model.Address{},
			Success: false,
		}
	}

	return IResponseAddress{
		Message: "Success updating address",
		Data:    address,
		Success: true,
	}
}

func (uc *AddressUseCase) DeleteAddress(id int) IResponseAddress {
	found, err := uc.repo.GetAddressByID(id)
	if err != nil {
		return IResponseAddress{
			Message: "Error getting address" + err.Error(),
			Data:    model.Address{},
			Success: false,
		}
	}

	_, err = uc.repo.DeleteAddress(id, found.UserID)
	if err != nil {
		return IResponseAddress{
			Message: "Error deleting address" + err.Error(),
			Data:    model.Address{},
			Success: false,
		}
	}

	return IResponseAddress{
		Message: "Success deleting address",
		Data:    found,
		Success: true,
	}
}
