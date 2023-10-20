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
	// Render swagger docs
	r.Handle("/swagger/*", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./docs"))))
	// Include routes
	router.SetupRoutes(r)
	// Start server
	http.ListenAndServe(":8000", r)
}
