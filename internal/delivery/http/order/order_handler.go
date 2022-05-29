package order_handler

import (
	"gokomodo/domain/repository"
	"gokomodo/domain/usecase"
	"gokomodo/internal/usecase/order"
	"gokomodo/pkg/service/jwt"

	"github.com/gorilla/mux"
)

type orderHandler struct {
	jwtService   jwt.JWTService
	orderUseCase usecase.OrderUseCase
}

func OrderHandler(
	r *mux.Router,
	jwtService jwt.JWTService,
	orderRepository repository.OrderRepository,
	productRepository repository.ProductRepository,
) {
	orderUseCase := order.NewOrderInteractor(orderRepository, productRepository)
	handler := &orderHandler{
		jwtService:   jwtService,
		orderUseCase: orderUseCase,
	}
	r.HandleFunc("/apis/v1/seller/order/list", handler.GetOrderSeller).Methods("GET")
	r.HandleFunc("/apis/v1/seller/order/{id}", handler.UpdateStatus).Methods("PUT")
	r.HandleFunc("/apis/v1/buyer/order/save", handler.Store).Methods("POST")
	r.HandleFunc("/apis/v1/buyer/order/list", handler.GetOrderBuyer).Methods("GET")
}
