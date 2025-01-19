package routes

import (
	"net/http"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/handlers/categories"
	"github.com/go-chi/chi/v5"
)

// CategoriesRoutes sets up the routes for category-related operations.
func CategoriesRoutes(db *database.Database) http.Handler {
    r := chi.NewRouter()

    // Define category-related routes
    r.Get("/", categories.HandleList(db))          // List all categories
    r.Get("/{id}", categories.HandleGet(db))       // Get a specific category by ID
    r.Post("/", categories.HandleCreate(db))       // Create a new category
    r.Put("/{id}", categories.HandleUpdate(db))    // Update a specific category by ID
    r.Delete("/{id}", categories.HandleDelete(db)) // Delete a specific category by ID

    return r
}
