package entity

import "time"

type Transaction struct {
	Id         int
	MerchantId int
	OutletId   int
	BillTotal  float64
	CreatedAt  time.Time
	CreatedBy  int
	UpdatedAt  time.Time
	UpdatedBy  int
}
