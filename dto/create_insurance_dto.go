package dto

import "github.com/shopspring/decimal"

type CreateInsuranceRequest struct {
	Name        string          `json:"name" binding:"required"`
	Description string          `json:"description" binding:"required"`
	Price       decimal.Decimal `json:"price" binding:"required"`
	Benefits    string          `json:"benefits" binding:"required"`
}
