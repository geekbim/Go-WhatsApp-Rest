package entity_test

import (
	"errors"
	"gokomodo/domain/entity"
	"gokomodo/domain/valueobject"
	"gokomodo/pkg/common"
	"gokomodo/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderDomain(t *testing.T) {
	orderDTO := testdata.NewOrderDTO()
	sellerId, _ := common.StringToID("35da70af-aa50-44dc-ae6b-060a0f9e6933")

	t.Run("NewOrder", func(t *testing.T) {
		res, err := entity.NewOrder(orderDTO)
		res.SetStatusAccepted()
		res.SetSeller(sellerId)
		res.SetTotalPrice(orderDTO.Qty, 10000)

		assert.Nil(t, err)
		assert.Equal(t, orderDTO.BuyerId, res.Buyer.Id)
		assert.Equal(t, orderDTO.ProductId, res.Product.Id)
		assert.Equal(t, orderDTO.Qty, res.Qty)
		assert.Equal(t, valueobject.ORDER_STATUS_ACCEPTED, res.Status.GetValue())
	})

	t.Run("NewOrderErrQty", func(t *testing.T) {
		orderDTO.Qty = 0
		err := errors.New("qty cannot be empty")
		expectedErr := []error{
			err,
		}

		res, errEntity := entity.NewOrder(orderDTO)

		assert.Equal(t, expectedErr, errEntity.Errors)
		assert.Nil(t, res)
	})
}
