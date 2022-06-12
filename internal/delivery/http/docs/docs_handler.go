package docs_handler

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func DocsHandler(r *mux.Router) {
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)
}
