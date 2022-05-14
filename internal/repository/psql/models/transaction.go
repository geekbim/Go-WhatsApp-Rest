package models

import "time"

type Transaction struct {
	Id           int       `dbq:"id"`
	MerchantName string    `dbq:"merchant_name"`
	OutletName   string    `dbq:"outlet_name"`
	Omzet        float64   `dbq:"omzet"`
	CreatedAt    time.Time `dbq:"created_at"`
	CreatedBy    int       `dbq:"created_by"`
	UpdatedAt    time.Time `dbq:"updated_at"`
	UpdatedBy    int       `dbq:"updated_by"`
}

func (Transaction) TableName() string {
	return "transactions"
}
