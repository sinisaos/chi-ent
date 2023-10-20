package main

import (
	"net/http"

	"github.com/sinisaos/chi-ent/router"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// Include routes
	router.SetupRoutes(r)
	// Start server
	http.ListenAndServe(":8000", r)
}
