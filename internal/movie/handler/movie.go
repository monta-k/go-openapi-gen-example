package handler

import (
	"log"
	"net/http"

	"github.com/monta-k/go-openapi-gen-example/internal/pkg/renderer"
	"github.com/monta-k/go-openapi-gen-example/openapi/movie/adapters"
)

type MovieHandler struct{}

func NewMovieHandler() *MovieHandler {
	return &MovieHandler{}
}

func (m *MovieHandler) GetV1Movies(w http.ResponseWriter, r *http.Request, params adapters.GetV1MoviesParams) {
	log.Printf("XUserID: %s", params.XUserID)
	if params.Title != nil {
		log.Printf("title: %s", *params.Title)
	}
	if params.Director != nil {
		log.Printf("Director: %s", *params.Director)
	}

	res := make([]*adapters.Movie, 0)
	res = append(res, &adapters.Movie{Id: "aaa", Title: "Babylon", Director: "Damien Chazelle"})
	res = append(res, &adapters.Movie{Id: "bbb", Title: "Pulp Fiction", Director: "Quentin Tarantino"})
	renderer.RenderJSON(w, http.StatusOK, res)
}
