package utils

type PaginationOption struct {
	Limit int
	Page  int
}

func NewPagination(page, limit int) PaginationOption {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	page = (page - 1) * limit
	return PaginationOption{
		Limit: limit,
		Page:  page,
	}
}
