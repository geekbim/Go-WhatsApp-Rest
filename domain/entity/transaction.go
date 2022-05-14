package entity

import "time"

type Transaction struct {
	Id        int
	Merchant  *Merchant
	Outlet    *Outlet
	Omzet     float64
	CreatedAt time.Time
	CreatedBy int
	UpdatedAt time.Time
	UpdatedBy int
}
