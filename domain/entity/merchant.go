package entity

import "time"

type Merchant struct {
	Id           int
	User         *User
	MerchantName string
	CreatedAt    time.Time
	CreatedBy    int
	UpdatedAt    time.Time
	UpdatedBy    int
}
