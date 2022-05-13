package models

import "time"

type User struct {
	Id        int       `dbq:"id"`
	Name      string    `dbq:"name"`
	UserName  string    `dbq:"user_name"`
	Password  string    `dbq:"password"`
	CreatedAt time.Time `dbq:"created_at"`
	CreatedBy int       `dbq:"created_by"`
	UpdatedAt time.Time `dbq:"updated_at"`
	UpdatedBy int       `dbq:"updated_by"`
}

func (User) TableName() string {
	return "users"
}
