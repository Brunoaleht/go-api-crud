package repository

import (
	"database/sql"
	"go-api-commerce/model"
	"log"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) *UserRepository {
	return &UserRepository{connection}
}

func (ur *UserRepository) GetUsers() ([]model.User, error) {
	query := "SELECT id, name, email, phone, group_id FROM user"
	rows, err := ur.connection.Query(query)
	if err != nil {
		log.Printf("Error querying user: %v", err)
		return []model.User{}, err
	}
	defer rows.Close()

	var userList []model.User
	var userObj model.User

	for rows.Next() {
		err := rows.Scan(&userObj.ID, &userObj.Name, &userObj.Email, &userObj.Phone, &userObj.GroupID)
		if err != nil {
			log.Printf("Error scanning user: %v", err)
			return []model.User{}, err
		}
		userList = append(userList, userObj)
	}

	return userList, nil
}

func (ur *UserRepository) CreateUser(user model.User) (int, error) {
	var id int
	query := "INSERT INTO user (name, email, password_hash, phone, group_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	result, err := ur.connection.Prepare(query)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	err = result.QueryRow(user.Name, user.Email, user.PasswordHash, user.Phone, user.GroupID).Scan(&id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer result.Close()

	return id, nil
}

func (ur *UserRepository) GetUserById(id int) (model.User, error) {
	query := "SELECT id, name, email, phone, group_id FROM user WHERE id = $1"
	row := ur.connection.QueryRow(query, id)

	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.GroupID)
	if err != nil {
		log.Println(err)
		return model.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) UpdateUser(user model.User) (int, error) {
	query := "UPDATE user SET name = $1, email = $2, phone = $3, group_id = $4 WHERE id = $5"
	result, err := ur.connection.Prepare(query)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	_, err = result.Exec(user.Name, user.Email, user.Phone, user.GroupID, user.ID)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer result.Close()

	return user.ID, nil
}

func (ur *UserRepository) DeleteUser(id int) (int, error) {
	query := "DELETE FROM user WHERE id = $1"
	result, err := ur.connection.Prepare(query)
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

func (ur *UserRepository) GetUserByEmail(email string) (model.User, error) {
	query := "SELECT id, name, email, phone, group_id FROM user WHERE email = $1"
	row := ur.connection.QueryRow(query, email)

	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.GroupID)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
