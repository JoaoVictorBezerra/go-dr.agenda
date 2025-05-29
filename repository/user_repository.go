package repository

import "database/sql"

type UserRepository struct {
	conn *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{
		conn: db,
	}
}
