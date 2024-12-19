package routes

import (
    "net/http"

    "github.com/CVWO/sample-go-app/internal/handlers/users"
    "github.com/go-chi/chi/v5"
)

// UsersRoutes sets up the routes for user-related operations.
func UsersRoutes() http.Handler {
    r := chi.NewRouter()

    // Define user-related routes
    r.Get("/", users.HandleList)          // List all users
    r.Get("/{username}", users.HandleGetByName)   // Get a specific user by name
    r.Get("/id/{id}", users.HandleGetByID)   // Get a specific user by ID
    r.Post("/", users.HandleCreate)       // Create a new user
    r.Put("/{username}", users.HandleUpdate) // Update a specific user by name
    r.Delete("/{username}", users.HandleDelete) // Delete a specific user by name

    return r
}
