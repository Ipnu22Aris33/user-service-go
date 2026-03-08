package server

import "net/http"

func SetupRouter(c *Container) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /user", c.UserHandler.CreateUser)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	return mux
}
