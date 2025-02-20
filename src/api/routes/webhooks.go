package routes

import (
	"knull/api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func WebhookRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/github", handlers.GitHubWebhook)

	return r
}
