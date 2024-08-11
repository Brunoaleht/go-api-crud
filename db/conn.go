package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
	host     = "go_db"
	port     = "5432"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbname, host, port)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to the database!")

	return db, nil
}
