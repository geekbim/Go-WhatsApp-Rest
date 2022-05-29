package response

import (
	"gokomodo/domain/entity"
	"time"
)

type Product struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ListProduct struct {
	Products []*Product `json:"products"`
	Count    int32      `json:"count"`
}

func MapProductDomainToResponse(product *entity.Product) *Product {
	return &Product{
		Id:          product.Id.String(),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func MapProductListDomainToResponse(products []*entity.Product, count int32) *ListProduct {
	res := make([]*Product, 0)

	for _, product := range products {
		res = append(res, MapProductDomainToResponse(product))
	}

	return &ListProduct{
		Products: res,
		Count:    count,
	}
}
