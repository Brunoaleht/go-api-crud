package usecase

import (
	"go-api-commerce/model"
	"go-api-commerce/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(pr repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		repository: pr,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	id, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}
	product.ID = id

	return product, nil
}

func (pu *ProductUseCase) GetProductByID(id int) (model.Product, error) {
	product, err := pu.repository.GetProductByID(id)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (pu *ProductUseCase) UpdateProduct(product model.Product) (model.Product, error) {
	err := pu.repository.UpdateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (pu *ProductUseCase) DeleteProduct(id int) error {
	err := pu.repository.DeleteProduct(id)
	if err != nil {
		return err
	}

	return nil
}
