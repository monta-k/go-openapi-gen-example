package movie

import (
	"fmt"
	"net/http"
	"os"

	"github.com/monta-k/go-openapi-gen-example/internal/pkg/renderer"
	"github.com/monta-k/go-openapi-gen-example/openapi/movie/adapters"
)

type Server struct {
	*http.Server
}

func NewServer() (*Server, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return nil, fmt.Errorf("PORT is not set")
	}

	deps := &Dependency{}
	deps.Inject()

	serverOptions := adapters.ChiServerOptions{
		ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			renderer.HandleError(r.Context(), w, err)
		},
	}
	handler := adapters.HandlerWithOptions(deps, serverOptions)

	server := http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
	return &Server{&server}, nil
}
