package mapper

import (
	"gokomodo/domain/entity"
	"gokomodo/domain/valueobject"
	"gokomodo/internal/repository/psql/models"
	"gokomodo/pkg/common"

	"github.com/rocketlaunchr/dbq/v2"
)

func ToDomainSellerOrder(m *models.OrderList) *entity.Order {
	id, _ := common.StringToID(m.Id)
	buyerId, _ := common.StringToID(m.UserId)
	productId, _ := common.StringToID(m.ProductId)
	status, _ := valueobject.NewOrderStatus(m.Status)
	order := &entity.Order{
		Id: id,
		Buyer: &entity.User{
			Id:      buyerId,
			Email:   m.UserEmail,
			Name:    m.UserName,
			Address: m.UserAddress,
		},
		Product: &entity.Product{
			Id:          productId,
			Name:        m.ProductName,
			Description: m.ProductDescription,
			Price:       m.ProductPrice,
			CreatedAt:   m.ProductCreatedAt,
			UpdatedAt:   m.ProductUpdatedAt,
		},
		Qty:        m.Qty,
		TotalPrice: m.TotalPrice,
		Status:     status,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}

	return order
}

func ToDomainSellerListOrder(models []*models.OrderList) []*entity.Order {
	domains := make([]*entity.Order, 0)

	for _, m := range models {
		d := ToDomainSellerOrder(m)
		domains = append(domains, d)
	}

	return domains
}

func ToDomainBuyerOrder(m *models.OrderList) *entity.Order {
	id, _ := common.StringToID(m.Id)
	sellerId, _ := common.StringToID(m.UserId)
	productId, _ := common.StringToID(m.ProductId)
	status, _ := valueobject.NewOrderStatus(m.Status)
	order := &entity.Order{
		Id: id,
		Seller: &entity.User{
			Id:      sellerId,
			Email:   m.UserEmail,
			Name:    m.UserName,
			Address: m.UserAddress,
		},
		Product: &entity.Product{
			Id:          productId,
			Name:        m.ProductName,
			Description: m.ProductDescription,
			Price:       m.ProductPrice,
			CreatedAt:   m.ProductCreatedAt,
			UpdatedAt:   m.ProductUpdatedAt,
		},
		Qty:        m.Qty,
		TotalPrice: m.TotalPrice,
		Status:     status,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}

	return order
}

func ToDomainBuyerListOrder(models []*models.OrderList) []*entity.Order {
	domains := make([]*entity.Order, 0)

	for _, m := range models {
		d := ToDomainBuyerOrder(m)
		domains = append(domains, d)
	}

	return domains
}

func ToModelOrder(d *entity.Order) *models.Order {
	order := &models.Order{
		Id:         d.Id.String(),
		BuyerId:    d.Buyer.Id.String(),
		SellerId:   d.Seller.Id.String(),
		ProductId:  d.Product.Id.String(),
		Qty:        d.Qty,
		TotalPrice: d.TotalPrice,
		Status:     d.Status.GetValue(),
		CreatedAt:  d.CreatedAt,
		UpdatedAt:  d.UpdatedAt,
	}

	return order
}

func ToModelListOrder(domains []*entity.Order) []*models.Order {
	models := make([]*models.Order, 0)

	for _, d := range domains {
		m := ToModelOrder(d)
		models = append(models, m)
	}

	return models
}

func DataDbqOrder(domain *entity.Order) []interface{} {
	return dbq.Struct(ToModelOrder(domain))
}

func ToDbqStructOrder(domain *entity.Order) (dbqStruct []interface{}) {
	dbqStruct = append(dbqStruct, DataDbqOrder(domain))
	return
}

func ToDbqStructOrderList(domains []*entity.Order) (dbqStruct []interface{}) {
	for _, domain := range domains {
		dbqData := DataDbqOrder(domain)
		dbqStruct = append(dbqStruct, dbqData)
	}
	return
}
