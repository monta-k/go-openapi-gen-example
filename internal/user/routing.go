package user

import "github.com/go-chi/chi/v5"

func routing(d *Dependency) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/ping", d.pingHandler.Ping)
	return r
}
