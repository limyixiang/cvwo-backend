package routes

import (
    "net/http"

    "github.com/CVWO/sample-go-app/internal/handlers/categories"
    "github.com/go-chi/chi/v5"
)

// CategoriesRoutes sets up the routes for category-related operations.
func CategoriesRoutes() http.Handler {
    r := chi.NewRouter()

    // Define category-related routes
    r.Get("/", categories.HandleList)          // List all categories
    r.Get("/{id}", categories.HandleGet)       // Get a specific category by ID
    r.Post("/", categories.HandleCreate)       // Create a new category
    r.Put("/{id}", categories.HandleUpdate)    // Update a specific category by ID
    r.Delete("/{id}", categories.HandleDelete) // Delete a specific category by ID

    return r
}
