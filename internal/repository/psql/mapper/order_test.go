package mapper_test

import (
	"gokomodo/domain/entity"
	"gokomodo/internal/repository/psql/mapper"
	"gokomodo/internal/repository/psql/models"
	"gokomodo/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderMapper(t *testing.T) {
	orderDTO := testdata.NewOrderDTO()
	orderDomain := testdata.NewOrder(orderDTO)
	ordersDomain := []*entity.Order{orderDomain}
	orderModel := testdata.NewOrderModel(orderDomain)
	ordersModel := []*models.Order{orderModel}
	orderListModel := testdata.NewOrderListModel()
	ordersListModel := []*models.OrderList{orderListModel}

	t.Run("ToDomainSellerListOrder", func(t *testing.T) {
		res := mapper.ToDomainSellerListOrder(ordersListModel)

		assert.Equal(t, res[0].Id, ordersDomain[0].Id)
		assert.Equal(t, res[0].Qty, ordersDomain[0].Qty)
		assert.Equal(t, res[0].TotalPrice, ordersDomain[0].TotalPrice)
		assert.Equal(t, res[0].Status, ordersDomain[0].Status)
	})

	t.Run("ToDomainBuyerListOrder", func(t *testing.T) {
		res := mapper.ToDomainBuyerListOrder(ordersListModel)

		assert.Equal(t, res[0].Id, ordersDomain[0].Id)
		assert.Equal(t, res[0].Qty, ordersDomain[0].Qty)
		assert.Equal(t, res[0].TotalPrice, ordersDomain[0].TotalPrice)
		assert.Equal(t, res[0].Status, ordersDomain[0].Status)
	})

	t.Run("ToModelListOrder", func(t *testing.T) {
		res := mapper.ToModelListOrder(ordersDomain)

		assert.Equal(t, res, ordersModel)
	})
}
