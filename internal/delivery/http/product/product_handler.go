package product_handler

import (
	"gokomodo/domain/repository"
	"gokomodo/domain/usecase"
	"gokomodo/internal/usecase/product"
	"gokomodo/pkg/service/jwt"

	"github.com/gorilla/mux"
)

type productHandler struct {
	jwtService     jwt.JWTService
	productUseCase usecase.ProductUseCase
}

func ProductHandler(
	r *mux.Router,
	jwtService jwt.JWTService,
	productRepository repository.ProductRepository,
) {
	productUseCase := product.NewProductInteractor(productRepository)
	handler := &productHandler{
		jwtService:     jwtService,
		productUseCase: productUseCase,
	}
	r.HandleFunc("/apis/v1/seller/product/list", handler.GetProductSeller).Methods("GET")
	r.HandleFunc("/apis/v1/seller/product/save", handler.Store).Methods("POST")
	r.HandleFunc("/apis/v1/buyer/product/list", handler.GetProductBuyer).Methods("GET")
}
