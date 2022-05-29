package models

import (
	"gokomodo/domain/valueobject"
	"time"
)

type Order struct {
	Id         string                      `dbq:"id"`
	BuyerId    string                      `dbq:"buyer_id"`
	SellerId   string                      `dbq:"seller_id"`
	ProductId  string                      `dbq:"product_id"`
	Qty        int                         `dbq:"qty"`
	TotalPrice int                         `dbq:"total_price"`
	Status     valueobject.OrderStatusEnum `dbq:"status"`
	CreatedAt  time.Time                   `dbq:"created_at"`
	UpdatedAt  time.Time                   `dbq:"updated_at"`
}

type OrderList struct {
	Id                 string                      `dbq:"id"`
	UserId             string                      `dbq:"user_id"`
	UserEmail          string                      `dbq:"user_email"`
	UserName           string                      `dbq:"user_name"`
	UserAddress        string                      `dbq:"user_address"`
	ProductId          string                      `dbq:"product_id"`
	ProductName        string                      `dbq:"product_name"`
	ProductDescription string                      `dbq:"product_description"`
	ProductPrice       int                         `dbq:"product_price"`
	ProductCreatedAt   time.Time                   `dbq:"product_created_at"`
	ProductUpdatedAt   time.Time                   `dbq:"product_updated_at"`
	Qty                int                         `dbq:"qty"`
	TotalPrice         int                         `dbq:"total_price"`
	Status             valueobject.OrderStatusEnum `dbq:"status"`
	CreatedAt          time.Time                   `dbq:"created_at"`
	UpdatedAt          time.Time                   `dbq:"updated_at"`
}

func (Order) TableName() string {
	return "orders"
}

func TableOrder() []string {
	return []string{
		"id",
		"buyer_id",
		"seller_id",
		"product_id",
		"qty",
		"total_price",
		"status",
		"created_at",
		"updated_at",
	}
}
