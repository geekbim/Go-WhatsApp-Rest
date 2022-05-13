package entity

import "time"

type User struct {
	Id        int
	Name      string
	UserName  string
	Password  string
	CreatedAt time.Time
	CreatedBy int
	UpdatedAt time.Time
	UpdatedBy int
}
