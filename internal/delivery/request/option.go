package request

import "go-rest-ddd/pkg/utils"

type Option struct {
	Pagination utils.PaginationOption
}

func NewOption(limit, page int) *Option {
	pagination := utils.NewPagination(page, limit)

	return &Option{
		Pagination: pagination,
	}
}
