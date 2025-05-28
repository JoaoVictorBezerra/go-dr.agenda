package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     string = "localhost"
	port     string = "5432"
	user     string = "admin"
	password string = "admin"
	dbName   string = "dr_agenda"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + dbName)

	return db, nil
}
