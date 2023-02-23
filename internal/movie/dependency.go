package movie

import "github.com/monta-k/go-openapi-gen-example/internal/movie/handler"

type Dependency struct {
	*handler.MovieHandler
}

func (d *Dependency) Inject() {
	d.MovieHandler = handler.NewMovieHandler()
}
