package usecase

import (
	"go-api-commerce/model"
	"go-api-commerce/repository"
)

type ProductUseCase struct {
	repository   repository.ProductRepository
	categoryRepo repository.CategoryRepository
}

type ProductListResponseApi struct {
	Message string          `json:"message"`
	Data    []model.Product `json:"data"`
	Success bool            `json:"success"`
}

type ProductResponseApi struct {
	Message string        `json:"message"`
	Data    model.Product `json:"data"`
	Success bool          `json:"success"`
}

func NewProductUseCase(repo repository.ProductRepository, categoryRepo repository.CategoryRepository) *ProductUseCase {
	return &ProductUseCase{
		repository:   repo,
		categoryRepo: categoryRepo,
	}
}

// func NewProductUseCase(pr repository.ProductRepository) *ProductUseCase {
// 	return &ProductUseCase{
// 		repository: pr,
// 	}
// }

func (pu *ProductUseCase) GetProducts() ProductListResponseApi {
	products, err := pu.repository.GetProducts()
	if err != nil {
		return ProductListResponseApi{
			Message: "Error getting products" + err.Error(),
			Data:    []model.Product{},
			Success: false,
		}
	}

	return ProductListResponseApi{
		Message: "Success getting products",
		Data:    products,
		Success: true,
	}

}

func (pu *ProductUseCase) CreateProduct(product model.Product) ProductResponseApi {
	id, err := pu.repository.CreateProduct(product)
	if err != nil {
		return ProductResponseApi{
			Message: "Error creating product" + err.Error(),
			Data:    model.Product{},
			Success: false,
		}
	}
	product.ID = id

	return ProductResponseApi{
		Message: "Success creating product",
		Data:    product,
		Success: true,
	}

}

func (pu *ProductUseCase) GetProductByID(id int) ProductResponseApi {
	product, err := pu.repository.GetProductByID(id)
	if err != nil {
		return ProductResponseApi{
			Message: "Error getting product" + err.Error(),
			Data:    model.Product{},
			Success: false,
		}
	}

	return ProductResponseApi{
		Message: "Success getting product",
		Data:    product,
		Success: true,
	}

}

func (pu *ProductUseCase) UpdateProduct(product model.Product) ProductResponseApi {
	_, err := pu.repository.GetProductByID(product.ID)
	if err != nil {
		return ProductResponseApi{
			Message: "Error not found product" + err.Error(),
			Data:    model.Product{},
			Success: false,
		}
	}

	id, err := pu.repository.UpdateProduct(product)
	if err != nil {
		return ProductResponseApi{
			Message: "Error updating product" + err.Error(),
			Data:    model.Product{},
			Success: false,
		}
	}
	product.ID = id

	return ProductResponseApi{
		Message: "Success updating product",
		Data:    product,
		Success: true,
	}
}

func (pu *ProductUseCase) DeleteProduct(id int) ProductResponseApi {
	productExist, err := pu.repository.GetProductByID(id)
	if err != nil {
		return ProductResponseApi{
			Message: "Error not found product" + err.Error(),
			Data:    model.Product{},
			Success: false,
		}
	}

	deletedId, err := pu.repository.DeleteProduct(id)
	if err != nil {
		return ProductResponseApi{
			Message: "Error deleting product" + err.Error(),
			Data:    model.Product{},
			Success: false,
		}
	}
	productExist.ID = deletedId

	return ProductResponseApi{
		Message: "Success deleting product",
		Data:    productExist,
		Success: true,
	}
}

func (pu *ProductUseCase) GetProductsByCategoryID(id int) ProductListResponseApi {
	_, err := pu.categoryRepo.GetCategoryByID(id)
	if err != nil {
		return ProductListResponseApi{
			Message: "Error not found category" + err.Error(),
			Data:    []model.Product{},
			Success: false,
		}
	}

	products, err := pu.repository.GetProductsByCategoryID(id)
	if err != nil {
		return ProductListResponseApi{
			Message: "Error getting products" + err.Error(),
			Data:    []model.Product{},
			Success: false,
		}
	}

	return ProductListResponseApi{
		Message: "Success getting products",
		Data:    products,
		Success: true,
	}
}
