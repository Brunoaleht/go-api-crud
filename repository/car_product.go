package repository

import (
	"database/sql"
	"go-api-commerce/model"
	"log"
)

type CarProductRepository struct {
	connection *sql.DB
}

func NewCarProductRepository(connection *sql.DB) *CarProductRepository {
	return &CarProductRepository{connection}
}

func (cpr *CarProductRepository) CreateCarProduct(carProduct model.CarProduct) (int, error) {
	var id int
	query := "INSERT INTO car_products (car_id, product_id, quantity, unit_price) VALUES ($1, $2, $3, $4) RETURNING id"
	result, err := cpr.connection.Prepare(query)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	err = result.QueryRow(carProduct.CarID, carProduct.ProductID, carProduct.Quantity, carProduct.UnitPrice).Scan(&id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer result.Close()

	return id, nil
}

func (cpr *CarProductRepository) GetCarProductByID(id int) (model.CarProduct, error) {
	query := "SELECT id, car_id, product_id, quantity, unit_price FROM car_products WHERE id = $1"
	row := cpr.connection.QueryRow(query, id)

	var carProduct model.CarProduct
	err := row.Scan(&carProduct.ID, &carProduct.CarID, &carProduct.ProductID, &carProduct.Quantity, &carProduct.UnitPrice)
	if err != nil {
		log.Printf("Error scanning car product: %v", err)
		return model.CarProduct{}, err
	}

	return carProduct, nil
}

func (cpr *CarProductRepository) UpdateCarProduct(carProduct model.CarProduct) (int, error) {
	query := `
		UPDATE car_products 
		SET 
			car_id = COALESCE(NULLIF($1::integer, NULL), car_id), 
			product_id = COALESCE(NULLIF($2::integer, NULL), product_id), 
			quantity = COALESCE(NULLIF($3::integer, NULL), quantity), 
			unit_price = COALESCE(NULLIF($4::numeric, NULL), unit_price) 
		WHERE id = $5`

	result, err := cpr.connection.Prepare(query)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	_, err = result.Exec(carProduct.CarID, carProduct.ProductID, carProduct.Quantity, carProduct.UnitPrice, carProduct.ID)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer result.Close()

	return carProduct.ID, nil
}

func (cpr *CarProductRepository) DeleteCarProduct(id int) (int, error) {
	query := "DELETE FROM car_products WHERE id = $1"
	result, err := cpr.connection.Prepare(query)
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

func (cpr *CarProductRepository) GetCarProductByProductID(productID int) (model.CarProduct, error) {
	query := "SELECT id, car_id, product_id, quantity, unit_price FROM car_products WHERE product_id = $1"
	row := cpr.connection.QueryRow(query, productID)

	var carProduct model.CarProduct
	err := row.Scan(&carProduct.ID, &carProduct.CarID, &carProduct.ProductID, &carProduct.Quantity, &carProduct.UnitPrice)
	if err != nil {
		log.Printf("Error scanning car product: %v", err)
		return model.CarProduct{}, err
	}

	return carProduct, nil
}

func (cpr *CarProductRepository) GetCarProductsByCarID(carID int) ([]model.CarProduct, error) {
	query := "SELECT id, car_id, product_id, quantity, unit_price FROM car_products WHERE car_id = $1"
	rows, err := cpr.connection.Query(query, carID)
	if err != nil {
		log.Printf("Error querying car product: %v", err)
		return []model.CarProduct{}, err
	}
	defer rows.Close()

	var carProductList []model.CarProduct
	var carProductObj model.CarProduct

	for rows.Next() {
		err := rows.Scan(&carProductObj.ID, &carProductObj.CarID, &carProductObj.ProductID, &carProductObj.Quantity, &carProductObj.UnitPrice)
		if err != nil {
			log.Printf("Error scanning car product: %v", err)
			return []model.CarProduct{}, err
		}
		carProductList = append(carProductList, carProductObj)
	}

	return carProductList, nil
}

func (cpr *CarProductRepository) GetCarProductsByCarIDWithTransaction(carID int, tx *Transaction) ([]model.CarProduct, error) {
	var carProducts []model.CarProduct

	rows, err := tx.tx.Query("SELECT id, car_id, product_id, quantity FROM car_products WHERE car_id = ?", carID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var carProduct model.CarProduct
		if err := rows.Scan(&carProduct.ID, &carProduct.CarID, &carProduct.ProductID, &carProduct.Quantity); err != nil {
			return nil, err
		}
		carProducts = append(carProducts, carProduct)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return carProducts, nil
}

func (cpr *CarProductRepository) CreateCarProductWithTransaction(carProduct model.CarProduct, tx *Transaction) (int, error) {
	result, err := tx.tx.Exec("INSERT INTO car_products (car_id, product_id, quantity) VALUES (?, ?, ?)", carProduct.CarID, carProduct.ProductID, carProduct.Quantity)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return int(id), nil
}
