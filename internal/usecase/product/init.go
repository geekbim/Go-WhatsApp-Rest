package product

import (
	"gokomodo/domain/repository"
	"gokomodo/domain/usecase"
)

type productInteractor struct {
	productRepository repository.ProductRepository
}

func NewProductInteractor(
	productRepository repository.ProductRepository,
) usecase.ProductUseCase {
	return &productInteractor{
		productRepository: productRepository,
	}
}
