package repository

import (
	"database/sql"
	"go-api-commerce/model"
	"log"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) *ProductRepository {
	return &ProductRepository{connection}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		log.Println(err)
		return []model.Product{}, err
	}
	defer rows.Close()

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err := rows.Scan(&productObj.ID, &productObj.Name, &productObj.Price)
		if err != nil {
			log.Println(err)
			return []model.Product{}, err
		}
		productList = append(productList, productObj)
	}

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int
	query := "INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id"
	result, err := pr.connection.Prepare(query)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	err = result.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer result.Close()

	return id, nil
}

func (pr *ProductRepository) GetProductByID(id int) (model.Product, error) {
	query := "SELECT id, product_name, price FROM product WHERE id = $1"
	row := pr.connection.QueryRow(query, id)

	var product model.Product
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		log.Println(err)
		return model.Product{}, err
	}

	return product, nil
}

func (pr *ProductRepository) UpdateProduct(product model.Product) error {
	query := "UPDATE product SET product_name = $1, price = $2 WHERE id = $3"
	_, err := pr.connection.Exec(query, product.Name, product.Price, product.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (pr *ProductRepository) DeleteProduct(id int) error {
	query := "DELETE FROM product WHERE id = $1"
	_, err := pr.connection.Exec(query, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
