package docs_handler

import (
	"go_wa_rest/internal/config"
	"go_wa_rest/valueobject"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func DocsHandler(r *mux.Router, cfg config.ServerConfig) {
	r.Handle("/doc.yaml", http.FileServer(http.Dir("./docs")))

	opts := middleware.SwaggerUIOpts{SpecURL: "/doc.yaml"}
	sh := middleware.SwaggerUI(opts, nil)

	env, err := valueobject.NewEnvTypeFromString(cfg.Env)
	if err != nil {
		panic(err)
	}
	if env.GetValue() == valueobject.Development {
		r.Handle("/docs", sh)
	}
}
