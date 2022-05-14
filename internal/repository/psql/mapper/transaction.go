package mapper

import (
	"majoo/domain/entity"
	"majoo/internal/repository/psql/models"
)

func ToDomainTransaction(m *models.Transaction) *entity.Transaction {
	transaction := &entity.Transaction{
		Id: m.Id,
		Merchant: &entity.Merchant{
			MerchantName: m.MerchantName,
		},
		Outlet: &entity.Outlet{
			OutletName: m.OutletName,
		},
		Omzet:     m.Omzet,
		CreatedAt: m.CreatedAt,
		CreatedBy: m.CreatedBy,
		UpdatedAt: m.UpdatedAt,
		UpdatedBy: m.UpdatedBy,
	}

	return transaction
}

func ToDomainListTransaction(models []*models.Transaction) []*entity.Transaction {
	domains := make([]*entity.Transaction, 0)

	for _, m := range models {
		d := ToDomainTransaction(m)
		domains = append(domains, d)
	}

	return domains
}

func ToModelTransaction(d *entity.Transaction) *models.Transaction {
	transaction := &models.Transaction{
		Id:           d.Id,
		MerchantName: d.Merchant.MerchantName,
		OutletName:   d.Outlet.OutletName,
		Omzet:        d.Omzet,
		CreatedAt:    d.CreatedAt,
		CreatedBy:    d.CreatedBy,
		UpdatedAt:    d.UpdatedAt,
		UpdatedBy:    d.UpdatedBy,
	}

	return transaction
}

func ToModelListTransaction(domains []*entity.Transaction) []*models.Transaction {
	models := make([]*models.Transaction, 0)

	for _, d := range domains {
		m := ToModelTransaction(d)
		models = append(models, m)
	}

	return models
}
