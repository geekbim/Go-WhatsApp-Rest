package entity

import "time"

type Outlet struct {
	Id         int
	MerchantId int
	OutletName string
	CreatedAt  time.Time
	CreatedBy  int
	UpdatedAt  time.Time
	UpdatedBy  int
}
