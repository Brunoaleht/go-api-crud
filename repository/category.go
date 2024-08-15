package repository

import (
	"database/sql"
	"go-api-commerce/model"
	"log"
)

type CategoryRepository struct {
	connection *sql.DB
}

func NewCategoryRepository(connection *sql.DB) *CategoryRepository {
	return &CategoryRepository{connection}
}

func (cr *CategoryRepository) GetCategories() ([]model.Category, error) {
	query := "SELECT id, name, description FROM category"
	rows, err := cr.connection.Query(query)
	if err != nil {
		log.Println(err)
		return []model.Category{}, err
	}
	defer rows.Close()

	var categoryList []model.Category
	var categoryObj model.Category

	for rows.Next() {
		err := rows.Scan(&categoryObj.ID, &categoryObj.Name, &categoryObj.Description)
		if err != nil {
			log.Println(err)
			return []model.Category{}, err
		}
		categoryList = append(categoryList, categoryObj)
	}

	return categoryList, nil
}

func (cr *CategoryRepository) CreateCategory(category model.Category) (int, error) {
	var id int
	query := "INSERT INTO category (name, description) VALUES ($1, $2) RETURNING id"
	result, err := cr.connection.Prepare(query)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	err = result.QueryRow(category.Name, category.Description).Scan(&id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer result.Close()

	return id, nil
}

func (cr *CategoryRepository) GetCategoryByID(id int) (model.Category, error) {
	query := "SELECT id, name, description FROM category WHERE id = $1"
	row := cr.connection.QueryRow(query, id)

	var category model.Category
	err := row.Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		log.Println(err)
		return model.Category{}, err
	}

	return category, nil
}

func (cr *CategoryRepository) UpdateCategory(category model.Category) (int, error) {
	query := "UPDATE category SET name = COALESCE(NULLIF($1, ''), name), description = COALESCE(NULLIF($2, ''), description) WHERE id = $3"
	result, err := cr.connection.Prepare(query)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	_, err = result.Exec(category.Name, category.Description, category.ID)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer result.Close()

	return category.ID, nil
}

func (cr *CategoryRepository) DeleteCategory(id int) (int, error) {
	query := "DELETE FROM category WHERE id = $1"
	result, err := cr.connection.Prepare(query)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	_, err = result.Exec(id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer result.Close()

	return id, nil
}

func (cr *CategoryRepository) GetCategoryByName(name string) (model.Category, error) {
	query := "SELECT id, name, description FROM category WHERE name = $1"
	row := cr.connection.QueryRow(query, name)

	var category model.Category
	err := row.Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		log.Println(err)
		return model.Category{}, err
	}

	return category, nil
}
