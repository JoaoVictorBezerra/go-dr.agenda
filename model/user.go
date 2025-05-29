package model

import "dr.agenda/enum"

type User struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Birthdate string    `json:"birthdate"`
	Role      enum.Role `json:"role"`
}
