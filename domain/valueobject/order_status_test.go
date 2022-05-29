package valueobject_test

import (
	"errors"
	"gokomodo/domain/valueobject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderStatus(t *testing.T) {
	t.Run("NewOrderStatusPendingFromInt", func(t *testing.T) {
		orderStatusEnum, err := valueobject.NewOrderStatusFromInt(0)

		assert.Nil(t, err)
		assert.Equal(t, valueobject.ORDER_STATUS_PENDING, orderStatusEnum)
	})

	t.Run("NewOrderStatusAcceptedFromInt", func(t *testing.T) {
		roleEnumFromInt, err := valueobject.NewOrderStatusFromInt(1)

		assert.Nil(t, err)
		assert.Equal(t, valueobject.ORDER_STATUS_ACCEPTED, roleEnumFromInt)
	})

	t.Run("NewOrderStatusFromIntErr", func(t *testing.T) {
		_, err := valueobject.NewOrderStatusFromInt(2)

		assert.Equal(t, errors.New("order status not found"), err)
	})

	t.Run("NewOrderStatusPendingFromString", func(t *testing.T) {
		orderStatusEnum, err := valueobject.NewOrderStatusFromString("PENDING")

		assert.Nil(t, err)
		assert.Equal(t, valueobject.ORDER_STATUS_PENDING, orderStatusEnum)
	})

	t.Run("NewOrderStatusAcceptedFromString", func(t *testing.T) {
		orderStatusEnum, err := valueobject.NewOrderStatusFromString("ACCEPTED")

		assert.Nil(t, err)
		assert.Equal(t, valueobject.ORDER_STATUS_ACCEPTED, orderStatusEnum)
	})

	t.Run("NewOrderStatusFromStringErr", func(t *testing.T) {
		_, err := valueobject.NewOrderStatusFromString("CANCEL")

		assert.Equal(t, errors.New("order status not found"), err)
	})

	t.Run("NewOrderStatusPending", func(t *testing.T) {
		const ORDER_STATUS_PENDING valueobject.OrderStatusEnum = 0
		orderStatus, err := valueobject.NewOrderStatus(ORDER_STATUS_PENDING)

		assert.Nil(t, err)
		assert.Equal(t, valueobject.ORDER_STATUS_PENDING, orderStatus.GetValue())
		assert.Equal(t, true, orderStatus.IsPending())
		assert.Equal(t, "PENDING", orderStatus.String())
	})

	t.Run("NewOrderStatusAccepted", func(t *testing.T) {
		const ORDER_STATUS_ACCEPTED valueobject.OrderStatusEnum = 1
		orderStatus, err := valueobject.NewOrderStatus(ORDER_STATUS_ACCEPTED)

		assert.Nil(t, err)
		assert.Equal(t, valueobject.ORDER_STATUS_ACCEPTED, orderStatus.GetValue())
		assert.Equal(t, true, orderStatus.IsAccepted())
		assert.Equal(t, "ACCEPTED", orderStatus.String())
	})

	t.Run("NewOrderStatusErr", func(t *testing.T) {
		const ORDER_STATUS_CANCEL valueobject.OrderStatusEnum = 2
		_, err := valueobject.NewOrderStatus(ORDER_STATUS_CANCEL)

		assert.Equal(t, errors.New("invalid order status, order status not supported"), err)
	})
}
