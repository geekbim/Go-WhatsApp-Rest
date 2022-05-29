package models

import "time"

type User struct {
	Id        string    `dbq:"id"`
	Email     string    `dbq:"email"`
	Name      string    `dbq:"name"`
	Password  string    `dbq:"password"`
	Address   string    `dbq:"address"`
	Role      string    `dbq:"role"`
	CreatedAt time.Time `dbq:"created_at"`
	UpdatedAt time.Time `dbq:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
