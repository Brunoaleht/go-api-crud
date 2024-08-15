package usecase

import (
	"go-api-commerce/model"
	"go-api-commerce/repository"
)

type CategoryListResponseApi struct {
	Message string           `json:"message"`
	Data    []model.Category `json:"data"`
	Success bool             `json:"success"`
}

type CategoryResponseApi struct {
	Message string         `json:"message"`
	Data    model.Category `json:"data"`
	Success bool           `json:"success"`
}

// CategoryUseCase is a struct that represents a category usecase
type CategoryUseCase struct {
	repository repository.CategoryRepository
}

// NewCategoryUseCase is a function to create a new category usecase
func NewCategoryUseCase(repo repository.CategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{
		repository: repo,
	}
}

func (cu *CategoryUseCase) GetCategories() CategoryListResponseApi {
	categories, err := cu.repository.GetCategories()
	if err != nil {
		return CategoryListResponseApi{
			Message: "Error getting categories" + err.Error(),
			Data:    []model.Category{},
			Success: false,
		}
	}

	return CategoryListResponseApi{
		Message: "Success getting categories",
		Data:    categories,
		Success: true,
	}
}

func (cu *CategoryUseCase) CreateCategory(category model.Category) CategoryResponseApi {
	categoryExists, err := cu.repository.GetCategoryByName(category.Name)
	if err != nil {
		return CategoryResponseApi{
			Message: "Error checking category existence" + err.Error(),
			Data:    model.Category{},
			Success: false,
		}
	}
	if categoryExists.ID != 0 {
		return CategoryResponseApi{
			Message: "Category already exists",
			Data:    model.Category{},
			Success: false,
		}
	}

	id, err := cu.repository.CreateCategory(category)
	if err != nil {
		return CategoryResponseApi{
			Message: "Error creating category" + err.Error(),
			Data:    model.Category{},
			Success: false,
		}
	}

	category.ID = id
	return CategoryResponseApi{
		Message: "Success creating category",
		Data:    category,
		Success: true,
	}
}

func (cu *CategoryUseCase) GetCategoryByID(id int) CategoryResponseApi {

	category, err := cu.repository.GetCategoryByID(id)
	if err != nil {
		return CategoryResponseApi{
			Message: "Error getting category" + err.Error(),
			Data:    model.Category{},
			Success: false,
		}
	}

	return CategoryResponseApi{
		Message: "Success getting category",
		Data:    category,
		Success: true,
	}
}

func (cu *CategoryUseCase) UpdateCategory(category model.Category) CategoryResponseApi {
	_, err := cu.repository.GetCategoryByID(category.ID)
	if err != nil {
		return CategoryResponseApi{
			Message: "Error not found category" + err.Error(),
			Data:    model.Category{},
			Success: false,
		}
	}

	id, err := cu.repository.UpdateCategory(category)
	if err != nil {
		return CategoryResponseApi{
			Message: "Error updating category" + err.Error(),
			Data:    model.Category{},
			Success: false,
		}
	}
	category.ID = id

	return CategoryResponseApi{
		Message: "Success updating category",
		Data:    category,
		Success: true,
	}
}

func (cu *CategoryUseCase) DeleteCategory(id int) CategoryResponseApi {
	categoryExist, err := cu.repository.GetCategoryByID(id)
	if err != nil {
		return CategoryResponseApi{
			Message: "Error not found category" + err.Error(),
			Data:    model.Category{},
			Success: false,
		}
	}

	_, err = cu.repository.DeleteCategory(id)
	if err != nil {
		return CategoryResponseApi{
			Message: "Error deleting category" + err.Error(),
			Data:    model.Category{},
			Success: false,
		}
	}

	return CategoryResponseApi{
		Message: "Success deleting category",
		Data:    categoryExist,
		Success: true,
	}
}
