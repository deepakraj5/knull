package routes

import "github.com/go-chi/chi/v5"

func BaseRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/accounts", func(r chi.Router) {
		r.Mount("/", AccountRoutes())
	})

	r.Route("/pipeline", func(r chi.Router) {
		r.Mount("/", PipelineRoutes())
	})

	r.Route("/webhook", func(r chi.Router) {
		r.Mount("/", WebhookRoutes())
	})

	return r
}
