package entity

import (
	"errors"
	"gokomodo/domain/valueobject"
	"gokomodo/pkg/common"
	"time"

	"github.com/hashicorp/go-multierror"
)

type Order struct {
	Id         common.ID
	Buyer      *User
	Seller     *User
	Product    *Product
	Qty        int
	TotalPrice int
	Status     *valueobject.OrderStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type OrderDTO struct {
	Id        *common.ID
	BuyerId   common.ID
	ProductId common.ID
	Qty       int
}

func NewOrder(orderDTO *OrderDTO) (*Order, *multierror.Error) {
	var multierr *multierror.Error

	if orderDTO.Id == nil {
		id := common.NewID()
		orderDTO.Id = &id
	}

	status, _ := valueobject.NewOrderStatus(valueobject.ORDER_STATUS_PENDING)

	order := &Order{
		Id: *orderDTO.Id,
		Buyer: &User{
			Id: orderDTO.BuyerId,
		},
		Product: &Product{
			Id: orderDTO.ProductId,
		},
		Qty:       orderDTO.Qty,
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if errValidate := order.Validate(); errValidate != nil {
		multierr = multierror.Append(multierr, errValidate)
	}

	if multierr != nil {
		return nil, multierr
	}

	return order, nil
}

func (order *Order) Validate() *multierror.Error {
	var multierr *multierror.Error

	if order.Qty == 0 {
		multierr = multierror.Append(multierr, errors.New("qty cannot be empty"))
	}

	return multierr
}

func (order *Order) SetStatusAccepted() {
	status, _ := valueobject.NewOrderStatus(valueobject.ORDER_STATUS_ACCEPTED)
	order.Status = status
}

func (order *Order) SetSeller(sellerId common.ID) {
	order.Seller = &User{Id: sellerId}
}

func (order *Order) SetTotalPrice(qty, price int) {
	order.TotalPrice = qty * price
}
