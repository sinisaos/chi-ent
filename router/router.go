package router

import (
	"github.com/sinisaos/chi-ent/handler"

	"github.com/go-chi/chi/v5"
)

// Setup api routes
func SetupRoutes(app chi.Router) {
	// Home route
	app.Get("/", handler.Index)
}
