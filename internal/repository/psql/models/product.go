package models

import "time"

type Product struct {
	Id          string    `dbq:"id"`
	Name        string    `dbq:"name"`
	Description string    `dbq:"description"`
	Price       int       `dbq:"price"`
	SellerId    string    `dbq:"seller_id"`
	CreatedAt   time.Time `dbq:"created_at"`
	UpdatedAt   time.Time `dbq:"updated_at"`
}

func (Product) TableName() string {
	return "products"
}

func TableProduct() []string {
	return []string{
		"id",
		"name",
		"description",
		"price",
		"seller_id",
		"created_at",
		"updated_at",
	}
}
