package mapper

import (
	"gokomodo/domain/entity"
	"gokomodo/internal/repository/psql/models"
	"gokomodo/pkg/common"

	"github.com/rocketlaunchr/dbq/v2"
)

func ToDomainProduct(m *models.Product) *entity.Product {
	id, _ := common.StringToID(m.Id)
	sellerId, _ := common.StringToID(m.SellerId)
	product := &entity.Product{
		Id:          id,
		Name:        m.Name,
		Description: m.Description,
		Price:       m.Price,
		Seller: &entity.User{
			Id: sellerId,
		},
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}

	return product
}

func ToDomainListProduct(models []*models.Product) []*entity.Product {
	domains := make([]*entity.Product, 0)

	for _, m := range models {
		d := ToDomainProduct(m)
		domains = append(domains, d)
	}

	return domains
}

func ToModelProduct(d *entity.Product) *models.Product {
	product := &models.Product{
		Id:          d.Id.String(),
		Name:        d.Name,
		Description: d.Description,
		Price:       d.Price,
		SellerId:    d.Seller.Id.String(),
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}

	return product
}

func ToModelListProduct(domains []*entity.Product) []*models.Product {
	models := make([]*models.Product, 0)

	for _, d := range domains {
		m := ToModelProduct(d)
		models = append(models, m)
	}

	return models
}

func DataDbqProduct(domain *entity.Product) []interface{} {
	return dbq.Struct(ToModelProduct(domain))
}

func ToDbqStructProduct(domain *entity.Product) (dbqStruct []interface{}) {
	dbqStruct = append(dbqStruct, DataDbqProduct(domain))
	return
}

func ToDbqStructProductList(domains []*entity.Product) (dbqStruct []interface{}) {
	for _, domain := range domains {
		dbqData := DataDbqProduct(domain)
		dbqStruct = append(dbqStruct, dbqData)
	}
	return
}
