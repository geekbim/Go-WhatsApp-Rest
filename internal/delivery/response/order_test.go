package response_test

import (
	"gokomodo/domain/entity"
	"gokomodo/internal/delivery/response"
	"gokomodo/testdata"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestOrderResponse(t *testing.T) {
	orderDTO := testdata.NewOrderDTO()
	order := testdata.NewOrder(orderDTO)
	orders := []*entity.Order{order}

	t.Run("MapSellerOrderListDomainToResponse", func(t *testing.T) {
		res := response.MapSellerOrderListDomainToResponse(orders, int32(len(orders)))

		assert.Equal(t, orders[0].Id.String(), res.Orders[0].Id)
	})

	t.Run("MapBuyerOrderListDomainToResponse", func(t *testing.T) {
		res := response.MapBuyerOrderListDomainToResponse(orders, int32(len(orders)))

		assert.Equal(t, orders[0].Id.String(), res.Orders[0].Id)
	})

	t.Run("MapOrderDomainToResponse", func(t *testing.T) {
		res := response.MapOrderDomainToResponse(order)

		assert.Equal(t, order.Id.String(), res.Id)
	})
}
