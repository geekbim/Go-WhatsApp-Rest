package mapper_test

import (
	"gokomodo/domain/entity"
	"gokomodo/internal/repository/psql/mapper"
	"gokomodo/internal/repository/psql/models"
	"gokomodo/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductMapper(t *testing.T) {
	productDTO := testdata.NewProductDTO()
	productDomain := testdata.NewProduct(productDTO)
	productsDomain := []*entity.Product{productDomain}
	productModel := testdata.NewProductModel(productDomain)
	productsModel := []*models.Product{productModel}

	t.Run("ToDomainListProduct", func(t *testing.T) {
		res := mapper.ToDomainListProduct(productsModel)

		assert.Equal(t, res, productsDomain)
	})

	t.Run("ToModelListProduct", func(t *testing.T) {
		res := mapper.ToModelListProduct(productsDomain)

		assert.Equal(t, res, productsModel)
	})
}
