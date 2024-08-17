package repository

import (
	"database/sql"
	"go-api-commerce/model"
	"log"
)

type AddressRepository struct {
	connection *sql.DB
}

func NewAddressRepository(connection *sql.DB) *AddressRepository {
	return &AddressRepository{connection}
}

func (ar *AddressRepository) GetAddressByUserID(userID int) ([]model.Address, error) {
	query := "SELECT id, user_id, street, city, state, zip_code, country FROM addresses WHERE user_id = $1"
	rows, err := ar.connection.Query(query, userID)
	if err != nil {
		log.Println(err)
		return []model.Address{}, err
	}
	defer rows.Close()

	var addressList []model.Address
	var addressObj model.Address

	for rows.Next() {
		err := rows.Scan(&addressObj.ID, &addressObj.UserID, &addressObj.Street, &addressObj.City, &addressObj.State, &addressObj.ZipCode, &addressObj.Country)
		if err != nil {
			log.Println(err)
			return []model.Address{}, err
		}
		addressList = append(addressList, addressObj)
	}

	return addressList, nil
}

func (ar *AddressRepository) CreateAddress(address model.Address) (int, error) {
	var id int
	query := "INSERT INTO addresses (user_id, street, city, state, zip_code, country) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	result, err := ar.connection.Prepare(query)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	err = result.QueryRow(address.UserID, address.Street, address.City, address.State, address.ZipCode, address.Country).Scan(&id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer result.Close()

	return id, nil
}

func (ar *AddressRepository) GetAddressByID(id int) (model.Address, error) {
	query := "SELECT id, user_id, street, city, state, zip_code, country FROM addresses WHERE id = $1"
	row := ar.connection.QueryRow(query, id)

	var address model.Address
	err := row.Scan(&address.ID, &address.UserID, &address.Street, &address.City, &address.State, &address.ZipCode, &address.Country)
	if err != nil {
		log.Println(err)
		return model.Address{}, err
	}

	return address, nil
}

func (ar *AddressRepository) UpdateAddress(address model.Address) error {
	query := "UPDATE addresses SET  street = COALESCE(NULLIF($1, ''), street), city = COALESCE(NULLIF($2, ''), city), state = COALESCE(NULLIF($3, ''), state), zip_code = COALESCE(NULLIF($4, ''), zip_code), country = COALESCE(NULLIF($5, ''), description) WHERE id = $6"
	_, err := ar.connection.Exec(query, address.Street, address.City, address.State, address.ZipCode, address.Country, address.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (ar *AddressRepository) DeleteAddress(id int, userID int) (int, error) {
	query := "DELETE FROM addresses WHERE id = $1 AND user_id = $2"
	result, err := ar.connection.Prepare(query)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer result.Close() // Fechar o statement após a execução

	_, err = result.Exec(id, userID)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return id, nil
}
