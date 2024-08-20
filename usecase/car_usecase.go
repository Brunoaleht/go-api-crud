package usecase

import (
	"go-api-commerce/model"
	"go-api-commerce/repository"
)

type IResponseCarWithProducts struct {
	Message string                `json:"message"`
	Data    model.CarWithProducts `json:"data"`
	Success bool                  `json:"success"`
}

type IResponseCars struct {
	Message string      `json:"message"`
	Data    []model.Car `json:"data"`
	Success bool        `json:"success"`
}

type IResponseCarProducts struct {
	Message string             `json:"message"`
	Data    []model.CarProduct `json:"data"`
	Success bool               `json:"success"`
}

type IResponseCar struct {
	Message string    `json:"message"`
	Data    model.Car `json:"data"`
	Success bool      `json:"success"`
}

type CarUseCase struct {
	cr  repository.CarRepository
	cpr repository.CarProductRepository
}

func NewCarUseCase(cr repository.CarRepository, cpr repository.CarProductRepository) *CarUseCase {
	return &CarUseCase{
		cr:  cr,
		cpr: cpr,
	}
}

func (uc *CarUseCase) GetCars() IResponseCars {
	cars, err := uc.cr.GetCars()
	if err != nil {
		return IResponseCars{
			Message: "Error getting cars" + err.Error(),
			Data:    []model.Car{},
			Success: false,
		}
	}

	return IResponseCars{
		Message: "Success getting cars",
		Data:    cars,
		Success: true,
	}
}

func (uc *CarUseCase) GetCarsByUserID(userID int) IResponseCars {
	cars, err := uc.cr.GetCarsByUserID(userID)
	if err != nil {
		return IResponseCars{
			Message: "Error getting cars" + err.Error(),
			Data:    []model.Car{},
			Success: false,
		}
	}

	return IResponseCars{
		Message: "Success getting cars",
		Data:    cars,
		Success: true,
	}
}

func (uc *CarUseCase) GetCarByID(id int) IResponseCar {
	car, err := uc.cr.GetCarByID(id)
	if err != nil {
		return IResponseCar{
			Message: "Error getting car" + err.Error(),
			Data:    model.Car{},
			Success: false,
		}
	}

	return IResponseCar{
		Message: "Success getting car",
		Data:    car,
		Success: true,
	}
}

func (uc *CarUseCase) GetCarWithProductsByCarID(carID int) IResponseCarWithProducts {
	carFound, err := uc.cr.GetCarByID(carID)
	if err != nil {
		return IResponseCarWithProducts{
			Message: "Error getting car" + err.Error(),
			Data: model.CarWithProducts{
				ID:       0,
				UserID:   0,
				Status:   model.CarStatusInactive,
				Products: []model.CarProduct{},
			},
			Success: false,
		}
	}

	carProducts, err := uc.cpr.GetCarProductsByCarID(carID)
	if err != nil {
		return IResponseCarWithProducts{
			Message: "Error getting product in car" + err.Error(),
			Data: model.CarWithProducts{
				ID:       carFound.ID,
				UserID:   carFound.UserID,
				Status:   carFound.Status,
				Products: []model.CarProduct{},
			},
			Success: false,
		}
	}

	return IResponseCarWithProducts{
		Message: "Success getting product in car",
		Data: model.CarWithProducts{
			ID:       carFound.ID,
			UserID:   carFound.UserID,
			Status:   carFound.Status,
			Products: carProducts,
		},
		Success: true,
	}
}

func (uc *CarUseCase) CreateCar(car model.Car) IResponseCar {
	id, err := uc.cr.CreateCar(car)
	if err != nil {
		return IResponseCar{
			Message: "Error creating car" + err.Error(),
			Data:    model.Car{},
			Success: false,
		}
	}

	car.ID = id
	return IResponseCar{
		Message: "Success creating car",
		Data:    car,
		Success: true,
	}
}

func (uc *CarUseCase) CreateCarWithProducts(car model.Car, carProducts []model.CarProduct) IResponseCarWithProducts {
	id, err := uc.cr.CreateCar(car)
	if err != nil {
		return IResponseCarWithProducts{
			Message: "Error creating car" + err.Error(),
			Data: model.CarWithProducts{
				ID:       0,
				UserID:   0,
				Status:   model.CarStatusInactive,
				Products: []model.CarProduct{},
			},
			Success: false,
		}
	}

	car.ID = id

	for i := 0; i < len(carProducts); i++ {
		carProducts[i].CarID = car.ID
		_, err := uc.cpr.CreateCarProduct(carProducts[i])
		if err != nil {
			return IResponseCarWithProducts{
				Message: "Error creating car product" + err.Error(),
				Data: model.CarWithProducts{
					ID:       car.ID,
					UserID:   car.UserID,
					Status:   car.Status,
					Products: []model.CarProduct{},
				},
				Success: false,
			}
		}
	}

	listCarProduct, err := uc.cpr.GetCarProductsByCarID(car.ID)
	if err != nil {
		return IResponseCarWithProducts{
			Message: "Error getting product in car" + err.Error(),
			Data: model.CarWithProducts{
				ID:       car.ID,
				UserID:   car.UserID,
				Status:   car.Status,
				Products: []model.CarProduct{},
			},
			Success: false,
		}
	}

	return IResponseCarWithProducts{
		Message: "Success creating car",
		Data: model.CarWithProducts{
			ID:       car.ID,
			UserID:   car.UserID,
			Status:   car.Status,
			Products: listCarProduct,
		},
		Success: true,
	}
}

func (uc *CarUseCase) AddProductToCar(userID, carID int, carProduct model.CarProduct) IResponseCarWithProducts {

	// Verifica se o carro existe
	carFound, err := uc.cr.GetCarByID(carID)
	if err != nil || carFound.ID == 0 {
		// Se o carrinho não for encontrado ou o ID do carrinho for zero, crie um novo carrinho
		newCar := model.Car{
			UserID: userID,
			Status: model.CarStatusInactive, // Ou o status que você deseja
		}

		carID, err := uc.cr.CreateCar(newCar)
		if err != nil {
			return IResponseCarWithProducts{
				Message: "Error creating new car: " + err.Error(),
				Data: model.CarWithProducts{
					ID:       0,
					UserID:   userID,
					Status:   "",
					Products: []model.CarProduct{},
				},
				Success: false,
			}
		}

		// Atualiza o ID do carro do carProduct com o novo carro criado
		carProduct.CarID = carID

		// Busca o novo carro criado
		carFound, err = uc.cr.GetCarByID(carID)
		if err != nil {
			return IResponseCarWithProducts{
				Message: "Error retrieving new car: " + err.Error(),
				Data: model.CarWithProducts{
					ID:       0,
					UserID:   userID,
					Status:   "",
					Products: []model.CarProduct{},
				},
				Success: false,
			}
		}
	}

	id, err := uc.cpr.CreateCarProduct(carProduct)
	if err != nil {
		return IResponseCarWithProducts{
			Message: "Error adding product to car" + err.Error(),
			Data: model.CarWithProducts{
				ID:       carFound.ID,
				UserID:   carFound.UserID,
				Status:   carFound.Status,
				Products: []model.CarProduct{},
			},
			Success: false,
		}
	}

	carProduct.ID = id

	listCarProduct, err := uc.cpr.GetCarProductsByCarID(carProduct.CarID)
	if err != nil {
		return IResponseCarWithProducts{
			Message: "Error getting product in car" + err.Error(),
			Data: model.CarWithProducts{
				ID:       carFound.ID,
				UserID:   carFound.UserID,
				Status:   carFound.Status,
				Products: []model.CarProduct{},
			},
			Success: false,
		}
	}

	return IResponseCarWithProducts{
		Message: "Success adding product to car",
		Data: model.CarWithProducts{
			ID:       carFound.ID,
			UserID:   carFound.UserID,
			Status:   carFound.Status,
			Products: listCarProduct,
		},
		Success: true,
	}
}

func (uc *CarUseCase) UpdateCarStatus(carID int, status string) IResponseCar {
	_, err := uc.cr.UpdateCarStatus(carID, status)
	if err != nil {
		return IResponseCar{
			Message: "Error updating car status" + err.Error(),
			Data:    model.Car{},
			Success: false,
		}
	}

	car, err := uc.cr.GetCarByID(carID)
	if err != nil {
		return IResponseCar{
			Message: "Error getting car" + err.Error(),
			Data:    model.Car{},
			Success: false,
		}
	}

	return IResponseCar{
		Message: "Success updating car status",
		Data:    car,
		Success: true,
	}
}

func (uc *CarUseCase) UpdateProductQuantity(userID, carID int, carProduct model.CarProduct) IResponseCarWithProducts {
	carFound, err := uc.cr.GetCarByID(carID)
	if err != nil {
		return IResponseCarWithProducts{
			Message: "Error getting car" + err.Error(),
			Data: model.CarWithProducts{
				ID:       0,
				UserID:   0,
				Status:   model.CarStatusInactive,
				Products: []model.CarProduct{},
			},
			Success: false,
		}
	}

	if carFound.UserID != userID {
		return IResponseCarWithProducts{
			Message: "Error getting not user's car",
			Data: model.CarWithProducts{
				ID:       carFound.ID,
				UserID:   userID,
				Status:   carFound.Status,
				Products: []model.CarProduct{},
			},
			Success: false,
		}
	}

	_, err = uc.cpr.UpdateCarProduct(carProduct)
	if err != nil {
		return IResponseCarWithProducts{
			Message: "Error updating product quantity" + err.Error(),
			Data: model.CarWithProducts{
				ID:       carFound.ID,
				UserID:   carFound.UserID,
				Status:   carFound.Status,
				Products: []model.CarProduct{},
			},
			Success: false,
		}
	}

	listCarProduct, err := uc.cpr.GetCarProductsByCarID(carProduct.CarID)
	if err != nil {
		return IResponseCarWithProducts{
			Message: "Error getting product in car" + err.Error(),
			Data: model.CarWithProducts{
				ID:       carFound.ID,
				UserID:   carFound.UserID,
				Status:   carFound.Status,
				Products: []model.CarProduct{},
			},
			Success: false,
		}
	}

	return IResponseCarWithProducts{
		Message: "Success updating product quantity",
		Data: model.CarWithProducts{
			ID:       carFound.ID,
			UserID:   carFound.UserID,
			Status:   carFound.Status,
			Products: listCarProduct,
		},
		Success: true,
	}
}

func (uc *CarUseCase) RemoveProductFromCar(userID, carID, productId int) IResponseCarWithProducts {

	carFound, err := uc.cr.GetCarByID(carID)
	if err != nil {
		return IResponseCarWithProducts{
			Message: "Error getting car" + err.Error(),
			Data: model.CarWithProducts{
				ID:       0,
				UserID:   0,
				Status:   model.CarStatusInactive,
				Products: []model.CarProduct{},
			},
			Success: false,
		}
	}
	if carFound.UserID != userID {
		return IResponseCarWithProducts{
			Message: "Error getting not user's car",
			Data: model.CarWithProducts{
				ID:       carFound.ID,
				UserID:   userID,
				Status:   carFound.Status,
				Products: []model.CarProduct{},
			},
			Success: false,
		}
	}

	_, err = uc.cpr.DeleteCarProduct(productId)
	if err != nil {
		return IResponseCarWithProducts{
			Message: "Error removing product from car" + err.Error(),
			Data: model.CarWithProducts{
				ID:       carFound.ID,
				UserID:   carFound.UserID,
				Status:   carFound.Status,
				Products: []model.CarProduct{},
			},
			Success: false,
		}
	}

	listCarProduct, err := uc.cpr.GetCarProductsByCarID(carID)
	if err != nil {
		return IResponseCarWithProducts{
			Message: "Error getting product in car" + err.Error(),
			Data: model.CarWithProducts{
				ID:       carFound.ID,
				UserID:   carFound.UserID,
				Status:   carFound.Status,
				Products: []model.CarProduct{},
			},
			Success: false,
		}
	}

	return IResponseCarWithProducts{
		Message: "Success removing product from car",
		Data: model.CarWithProducts{
			ID:       carFound.ID,
			UserID:   carFound.UserID,
			Status:   carFound.Status,
			Products: listCarProduct,
		},
		Success: true,
	}
}
