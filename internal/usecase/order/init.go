package order

import (
	"gokomodo/domain/repository"
	"gokomodo/domain/usecase"
)

type orderInteractor struct {
	orderRepository   repository.OrderRepository
	productRepository repository.ProductRepository
}

func NewOrderInteractor(
	orderRepository repository.OrderRepository,
	productRepository repository.ProductRepository,
) usecase.OrderUseCase {
	return &orderInteractor{
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
}
