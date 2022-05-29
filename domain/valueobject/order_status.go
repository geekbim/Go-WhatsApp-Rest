package valueobject

import (
	"errors"
)

type OrderStatusEnum int

const (
	ORDER_STATUS_PENDING  OrderStatusEnum = 0
	ORDER_STATUS_ACCEPTED OrderStatusEnum = 1
)

type OrderStatus struct {
	value OrderStatusEnum
}

func (o *OrderStatus) String() string {
	orderStatusString := ""

	switch o.value {
	case ORDER_STATUS_PENDING:
		orderStatusString = "PENDING"
	case ORDER_STATUS_ACCEPTED:
		orderStatusString = "ACCEPTED"
	}

	return orderStatusString
}

func (o *OrderStatus) GetValue() OrderStatusEnum {
	return o.value
}

func (o *OrderStatus) IsPending() bool {
	return o.value == ORDER_STATUS_PENDING
}

func (o *OrderStatus) IsAccepted() bool {
	return o.value == ORDER_STATUS_ACCEPTED
}

func NewOrderStatus(value OrderStatusEnum) (*OrderStatus, error) {
	if value < 0 || value > 1 {
		return nil, errors.New("invalid order status, order status not supported")
	}

	return &OrderStatus{value: value}, nil
}

func NewOrderStatusFromString(str string) (OrderStatusEnum, error) {
	var value OrderStatusEnum

	switch str {
	case "PENDING":
		value = ORDER_STATUS_PENDING
	case "ACCEPTED":
		value = ORDER_STATUS_ACCEPTED
	default:
		return value, errors.New("order status not found")
	}

	return value, nil
}

func NewOrderStatusFromInt(num int) (OrderStatusEnum, error) {
	var value OrderStatusEnum

	switch num {
	case 0:
		value = ORDER_STATUS_PENDING
	case 1:
		value = ORDER_STATUS_ACCEPTED
	default:
		return value, errors.New("order status not found")
	}

	return value, nil
}
