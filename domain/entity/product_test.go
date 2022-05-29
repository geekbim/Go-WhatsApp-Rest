package entity_test

import (
	"errors"
	"gokomodo/domain/entity"
	"gokomodo/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductDomain(t *testing.T) {
	productDTO := testdata.NewProductDTO()

	t.Run("NewProduct", func(t *testing.T) {
		res, err := entity.NewProduct(productDTO)

		assert.Nil(t, err)
		assert.Equal(t, productDTO.Name, res.Name)
		assert.Equal(t, productDTO.Description, res.Description)
		assert.Equal(t, productDTO.Price, res.Price)
	})

	t.Run("NewProductErrName", func(t *testing.T) {
		productDTO.Name = ""
		err := errors.New("product name cannot be empty")
		expectedErr := []error{
			err,
		}

		res, errEntity := entity.NewProduct(productDTO)

		assert.Equal(t, expectedErr, errEntity.Errors)
		assert.Nil(t, res)
	})

	t.Run("NewProductErrDescription", func(t *testing.T) {
		productDTO.Name = "Tolak Angin"
		productDTO.Description = ""
		err := errors.New("product description cannot be empty")
		expectedErr := []error{
			err,
		}

		res, errEntity := entity.NewProduct(productDTO)

		assert.Equal(t, expectedErr, errEntity.Errors)
		assert.Nil(t, res)
	})

	t.Run("NewProductErrPrice", func(t *testing.T) {
		productDTO.Name = "Tolak Angin"
		productDTO.Description = "Herbal masuk angin"
		productDTO.Price = 0
		err := errors.New("product price cannot be empty")
		expectedErr := []error{
			err,
		}

		res, errEntity := entity.NewProduct(productDTO)

		assert.Equal(t, expectedErr, errEntity.Errors)
		assert.Nil(t, res)
	})
}
