package routes

import "github.com/go-chi/chi/v5"

func BaseRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/accounts", func(r chi.Router) {
		r.Mount("/", AccountRoutes())
	})

	return r
}
