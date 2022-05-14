package testdata

import (
	"majoo/domain/entity"
	"time"
)

func NewTransaction() *entity.Transaction {
	return &entity.Transaction{
		Id: 1,
		Merchant: &entity.Merchant{
			MerchantName: "merchant name",
		},
		Outlet: &entity.Outlet{
			OutletName: "outlet name",
		},
		CreatedAt: time.Time{},
		CreatedBy: 1,
		UpdatedAt: time.Time{},
		UpdatedBy: 1,
	}
}
