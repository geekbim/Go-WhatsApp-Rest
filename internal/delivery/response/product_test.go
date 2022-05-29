package response_test

import (
	"gokomodo/domain/entity"
	"gokomodo/internal/delivery/response"
	"gokomodo/testdata"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestProductResponse(t *testing.T) {
	productDTO := testdata.NewProductDTO()
	product := testdata.NewProduct(productDTO)
	products := []*entity.Product{product}

	res := response.MapProductListDomainToResponse(products, int32(len(products)))

	assert.Equal(t, products[0].Id.String(), res.Products[0].Id)
	assert.Equal(t, products[0].Name, res.Products[0].Name)
	assert.Equal(t, products[0].Description, res.Products[0].Description)
	assert.Equal(t, products[0].Price, res.Products[0].Price)
}
