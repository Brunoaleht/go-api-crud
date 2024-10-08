package repository

import (
	"database/sql"
	"go-api-commerce/model"
	"log"
)

type CarRepository struct {
	connection *sql.DB
}

func NewCarRepository(connection *sql.DB) *CarRepository {
	return &CarRepository{connection}
}

func (cr *CarRepository) GetCars() ([]model.Car, error) {
	query := "SELECT id, user_id, status, created_at, updated_at FROM cars"
	rows, err := cr.connection.Query(query)
	if err != nil {
		log.Printf("Error querying car: %v", err)
		return []model.Car{}, err
	}
	defer rows.Close()

	var carList []model.Car
	var carObj model.Car

	for rows.Next() {
		err := rows.Scan(&carObj.ID, &carObj.UserID, &carObj.Status, &carObj.CreatedAt, &carObj.UpdatedAt)
		if err != nil {
			log.Printf("Error scanning car: %v", err)
			return []model.Car{}, err
		}
		carList = append(carList, carObj)
	}

	return carList, nil
}

func (cr *CarRepository) CreateCar(car model.Car) (int, error) {
	var id int
	query := "INSERT INTO cars (user_id, status, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id"
	result, err := cr.connection.Prepare(query)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	err = result.QueryRow(car.UserID, car.Status, car.CreatedAt, car.UpdatedAt).Scan(&id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer result.Close()

	return id, nil
}

func (cr *CarRepository) GetCarByID(id int) (model.Car, error) {
	query := "SELECT id, user_id, status, created_at, updated_at FROM cars WHERE id = $1"
	row := cr.connection.QueryRow(query, id)

	var carObj model.Car
	err := row.Scan(&carObj.ID, &carObj.UserID, &carObj.Status, &carObj.CreatedAt, &carObj.UpdatedAt)
	if err != nil {
		log.Printf("Error scanning car: %v", err)
		return model.Car{}, err
	}

	return carObj, nil
}

func (cr *CarRepository) GetCarsByUserID(userID int) ([]model.Car, error) {
	query := "SELECT id, user_id, status, created_at, updated_at FROM cars WHERE user_id = $1"
	rows, err := cr.connection.Query(query, userID)
	if err != nil {
		log.Println(err)
		return []model.Car{}, err
	}
	defer rows.Close()

	var carList []model.Car
	var carObj model.Car

	for rows.Next() {
		err := rows.Scan(&carObj.ID, &carObj.UserID, &carObj.Status, &carObj.CreatedAt, &carObj.UpdatedAt)
		if err != nil {
			log.Println(err)
			return []model.Car{}, err
		}
		carList = append(carList, carObj)
	}

	return carList, nil
}

func (cr *CarRepository) GetCarActiveByUserID(userID int) (model.Car, error) {
	query := "SELECT id, user_id, status, created_at, updated_at FROM cars WHERE user_id = $1 AND status = 'active'"
	row := cr.connection.QueryRow(query, userID)

	var carObj model.Car
	err := row.Scan(&carObj.ID, &carObj.UserID, &carObj.Status, &carObj.CreatedAt, &carObj.UpdatedAt)
	if err != nil {
		log.Println(err)
		return model.Car{}, err
	}

	return carObj, nil
}

func (cr *CarRepository) UpdateCarStatus(carID int, status string) (int, error) {
	query := "UPDATE cars SET status = $1 WHERE id = $2"
	result, err := cr.connection.Prepare(query)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer result.Close()

	_, err = result.Exec(status, carID)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return carID, nil
}

func (cr *CarRepository) DeleteCar(id int) (int, error) {
	query := "DELETE FROM cars WHERE id = $1"
	result, err := cr.connection.Prepare(query)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer result.Close()

	_, err = result.Exec(id)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return id, nil
}

// Transaction
type Transaction struct {
	tx *sql.Tx
}

func (t *Transaction) Commit() error {
	return t.tx.Commit()
}

func (t *Transaction) Rollback() error {
	return t.tx.Rollback()
}

func (cr *CarRepository) BeginTransaction() (*Transaction, error) {
	tx, err := cr.connection.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &Transaction{tx: tx}, nil
}

func (cr *CarRepository) GetCarActiveByUserIDWithTransaction(userID int, tx *Transaction) (model.Car, error) {
	var car model.Car
	err := tx.tx.QueryRow("SELECT id, user_id, status FROM cars WHERE user_id = $1 AND status = $2", userID, model.CarStatusActive).Scan(&car.ID, &car.UserID, &car.Status)
	if err != nil {
		log.Println(err)
		return car, err
	}
	return car, nil
}

func (cr *CarRepository) CreateCarWithTransaction(car model.Car, tx *Transaction) (int, error) {
	result, err := tx.tx.Exec("INSERT INTO cars (user_id, status) VALUES ($1, $2)", car.UserID, car.Status)
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
