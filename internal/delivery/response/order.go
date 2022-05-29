package response

import (
	"gokomodo/domain/entity"
	"time"
)

type Order struct {
	Id         string    `json:"id"`
	BuyerId    string    `json:"buyerId"`
	SellerId   string    `json:"sellerId"`
	ProductId  string    `json:"productId"`
	Qty        int       `json:"qty"`
	TotalPrice int       `json:"totalPrice"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type OrderList struct {
	Id         string     `json:"id"`
	Buyer      *UserOrder `json:"buyer"`
	Seller     *UserOrder `json:"seller"`
	Product    *Product   `json:"product"`
	Qty        int        `json:"qty"`
	TotalPrice int        `json:"totalPrice"`
	Status     string     `json:"status"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
}

type UserOrder struct {
	Id      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type ListOrder struct {
	Orders []*OrderList `json:"orders"`
	Count  int32        `json:"count"`
}

func MapOrderDomainToResponse(order *entity.Order) *Order {
	return &Order{
		Id:         order.Id.String(),
		BuyerId:    order.Buyer.Id.String(),
		SellerId:   order.Seller.Id.String(),
		ProductId:  order.Product.Id.String(),
		Qty:        order.Qty,
		TotalPrice: order.TotalPrice,
		Status:     order.Status.String(),
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
	}
}

func MapSellerOrderDomainToResponseList(order *entity.Order) *OrderList {
	return &OrderList{
		Id: order.Id.String(),
		Buyer: &UserOrder{
			Id:      order.Buyer.Id.String(),
			Email:   order.Buyer.Email,
			Name:    order.Buyer.Name,
			Address: order.Buyer.Address,
		},
		Product: &Product{
			Id:          order.Product.Id.String(),
			Name:        order.Product.Name,
			Description: order.Product.Description,
			Price:       order.Product.Price,
			CreatedAt:   order.Product.CreatedAt,
			UpdatedAt:   order.Product.UpdatedAt,
		},
		Qty:        order.Qty,
		TotalPrice: order.TotalPrice,
		Status:     order.Status.String(),
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
	}
}

func MapSellerOrderListDomainToResponse(orders []*entity.Order, count int32) *ListOrder {
	res := make([]*OrderList, 0)

	for _, order := range orders {
		res = append(res, MapSellerOrderDomainToResponseList(order))
	}

	return &ListOrder{
		Orders: res,
		Count:  count,
	}
}

func MapBuyerOrderDomainToResponseList(order *entity.Order) *OrderList {
	return &OrderList{
		Id: order.Id.String(),
		Seller: &UserOrder{
			Id:      order.Seller.Id.String(),
			Email:   order.Seller.Email,
			Name:    order.Seller.Name,
			Address: order.Seller.Address,
		},
		Product: &Product{
			Id:          order.Product.Id.String(),
			Name:        order.Product.Name,
			Description: order.Product.Description,
			Price:       order.Product.Price,
			CreatedAt:   order.Product.CreatedAt,
			UpdatedAt:   order.Product.UpdatedAt,
		},
		Qty:        order.Qty,
		TotalPrice: order.TotalPrice,
		Status:     order.Status.String(),
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
	}
}

func MapBuyerOrderListDomainToResponse(orders []*entity.Order, count int32) *ListOrder {
	res := make([]*OrderList, 0)

	for _, order := range orders {
		res = append(res, MapBuyerOrderDomainToResponseList(order))
	}

	return &ListOrder{
		Orders: res,
		Count:  count,
	}
}
