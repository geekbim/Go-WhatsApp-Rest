package testdata

import (
	"gokomodo/domain/entity"
	"gokomodo/domain/valueobject"
	"gokomodo/internal/repository/psql/mapper"
	"gokomodo/internal/repository/psql/models"
	"gokomodo/pkg/common"
	"time"
)

func NewOrderDTO() *entity.OrderDTO {
	buyerId, _ := common.StringToID("3b584960-f83f-480f-b2ee-2f66e0e4ac69")
	productId, _ := common.StringToID("bf1796fe-cc7c-4a46-a422-da7317c8916f")
	return &entity.OrderDTO{
		Id:        nil,
		BuyerId:   buyerId,
		ProductId: productId,
		Qty:       10,
	}
}

func NewOrder(orderDTO *entity.OrderDTO) *entity.Order {
	id, _ := common.StringToID("328f1db5-a1a8-42e5-a4c2-a1d465804160")
	sellerId, _ := common.StringToID("35da70af-aa50-44dc-ae6b-060a0f9e6933")
	status, _ := valueobject.NewOrderStatus(valueobject.ORDER_STATUS_PENDING)
	return &entity.Order{
		Id: id,
		Buyer: &entity.User{
			Id: orderDTO.BuyerId,
		},
		Seller: &entity.User{
			Id: sellerId,
		},
		Product: &entity.Product{
			Id: orderDTO.ProductId,
		},
		Qty:        orderDTO.Qty,
		TotalPrice: 100000,
		Status:     status,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}
}

func NewOrderModel(order *entity.Order) *models.Order {
	return mapper.ToModelOrder(order)
}

func NewOrderListModel() *models.OrderList {
	status, _ := valueobject.NewOrderStatus(valueobject.ORDER_STATUS_PENDING)
	return &models.OrderList{
		Id:                 "328f1db5-a1a8-42e5-a4c2-a1d465804160",
		UserId:             "35da70af-aa50-44dc-ae6b-060a0f9e6933",
		UserEmail:          "seller@email.com",
		UserName:           "Seller",
		UserAddress:        "Jakarta Barat",
		ProductId:          "bf1796fe-cc7c-4a46-a422-da7317c8916f",
		ProductName:        "Tolak Angin",
		ProductDescription: "Herbal masuk angin",
		ProductPrice:       10000,
		ProductCreatedAt:   time.Time{},
		ProductUpdatedAt:   time.Time{},
		Qty:                10,
		TotalPrice:         100000,
		Status:             status.GetValue(),
		CreatedAt:          time.Time{},
		UpdatedAt:          time.Time{},
	}
}
