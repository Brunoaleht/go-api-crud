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
	query := "SELECT id, name, price, description, category_id, stock_quantity FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		log.Println(err)
		return []model.Product{}, err
	}
	defer rows.Close()

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err := rows.Scan(&productObj.ID, &productObj.Name, &productObj.Price, &productObj.Description, &productObj.CategoryID, &productObj.StockQuantity)
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
	query := "INSERT INTO product (name, description, category_id, price, stock_quantity ) VALUES ($1, $2, $3, $4, $5 ) RETURNING id"
	result, err := pr.connection.Prepare(query)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	err = result.QueryRow(product.Name, product.Description, product.CategoryID, product.Price, product.StockQuantity).Scan(&id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer result.Close()

	return id, nil
}

func (pr *ProductRepository) GetProductByID(id int) (model.Product, error) {
	query := "SELECT id, name, price, description, category_id, stock_quantity, created_at, updated_at FROM product WHERE id = $1"
	row := pr.connection.QueryRow(query, id)

	var product model.Product
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.CategoryID, &product.StockQuantity, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		log.Println(err)
		return model.Product{}, err
	}

	return product, nil
}

func (pr *ProductRepository) UpdateProduct(product model.Product) (int, error) {
	query := "UPDATE product SET name = COALESCE(NULLIF($1, ''), name), description = COALESCE(NULLIF($2, ''), description), price = COALESCE(NULLIF($3, ''), price), category_id =COALESCE(NULLIF($4, ''), category_id), stock_quantity = COALESCE(NULLIF($5, ''), stock_quantity) WHERE id = $6"
	// query := "UPDATE product SET product_name = $1, price = $2 WHERE id = $3"
	_, err := pr.connection.Exec(query, product.Name, product.Description, product.Price, product.CategoryID, product.StockQuantity, product.ID)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return product.ID, nil
}

func (pr *ProductRepository) DeleteProduct(id int) (int, error) {
	query := "DELETE FROM product WHERE id = $1"
	result, err := pr.connection.Prepare(query)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	_, err = result.Exec(id)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return id, nil
}

func (pr *ProductRepository) GetProductsByCategoryID(id int) ([]model.Product, error) {
	query := "SELECT id, name, price, description, category_id, stock_quantity FROM product WHERE category_id = $1"
	rows, err := pr.connection.Query(query, id)
	if err != nil {
		log.Println(err)
		return []model.Product{}, err
	}
	defer rows.Close()

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err := rows.Scan(&productObj.ID, &productObj.Name, &productObj.Price, &productObj.Description, &productObj.CategoryID, &productObj.StockQuantity)
		if err != nil {
			log.Println(err)
			return []model.Product{}, err
		}
		productList = append(productList, productObj)
	}

	return productList, nil
}
