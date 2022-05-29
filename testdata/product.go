package testdata

import (
	"gokomodo/domain/entity"
	"gokomodo/internal/repository/psql/mapper"
	"gokomodo/internal/repository/psql/models"
	"gokomodo/pkg/common"
	"time"
)

func NewProductDTO() *entity.ProductDTO {
	sellerId, _ := common.StringToID("35da70af-aa50-44dc-ae6b-060a0f9e6933")
	return &entity.ProductDTO{
		Id:          nil,
		Name:        "Tolak Angin",
		Description: "Herbal masuk angin",
		Price:       10000,
		SellerId:    sellerId,
	}
}

func NewProduct(productDTO *entity.ProductDTO) *entity.Product {
	id, _ := common.StringToID("bf1796fe-cc7c-4a46-a422-da7317c8916f")
	return &entity.Product{
		Id:          id,
		Name:        productDTO.Name,
		Description: productDTO.Description,
		Price:       productDTO.Price,
		Seller: &entity.User{
			Id: productDTO.SellerId,
		},
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func NewProductModel(product *entity.Product) *models.Product {
	return mapper.ToModelProduct(product)
}
