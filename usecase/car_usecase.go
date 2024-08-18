package usecase

import "go-api-commerce/repository"

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
