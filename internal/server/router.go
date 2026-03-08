package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func SetupRouter(c *Container) http.Handler {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", c.UserHandler.CreateUser)
	})

	return r
}
