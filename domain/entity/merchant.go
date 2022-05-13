package entity

import "time"

type Merchant struct {
	Id           int
	UserId       int
	MerchantName string
	CreatedAt    time.Time
	CreatedBy    int
	UpdatedAt    time.Time
	UpdatedBy    int
}
