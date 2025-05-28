package model

import (
	"dr.agenda/enum"
)

type HealthInsurance struct {
	Id          int                        `json:"id"`
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Price       float32                    `json:"price"`
	Benefits    string                     `json:"benefits"`
	Status      enum.HealthInsuranceStatus `json:"status"`
}
