package response

import (
	"majoo/domain/entity"
	"time"
)

type Transaction struct {
	Id           int       `json:"id"`
	MerchantName string    `json:"merchantName"`
	OutletName   string    `json:"outletName"`
	Omzet        float64   `json:"omzet"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type ListTransaction struct {
	Transactions []*Transaction `json:"transactions"`
	Count        int32          `json:"count"`
}

func MapTransactionDomainToResponse(transaction *entity.Transaction) *Transaction {
	return &Transaction{
		Id:           transaction.Id,
		MerchantName: transaction.Merchant.MerchantName,
		OutletName:   transaction.Outlet.OutletName,
		Omzet:        transaction.Omzet,
		CreatedAt:    transaction.CreatedAt,
		UpdatedAt:    transaction.UpdatedAt,
	}
}

func MapTransactionListDomainToResponse(transactions []*entity.Transaction, count int32) *ListTransaction {
	res := make([]*Transaction, 0)

	for _, transaction := range transactions {
		res = append(res, MapTransactionDomainToResponse(transaction))
	}

	return &ListTransaction{
		Transactions: res,
		Count:        count,
	}
}
