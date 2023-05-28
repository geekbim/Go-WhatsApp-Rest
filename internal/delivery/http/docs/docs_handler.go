package docs_handler

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func DocsHandler(r *mux.Router) {
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)
}
