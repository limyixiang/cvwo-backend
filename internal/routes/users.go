package routes

import (
	"net/http"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/handlers/users"
	"github.com/go-chi/chi/v5"
)

// UsersRoutes sets up the routes for user-related operations.
func UsersRoutes(db *database.Database) http.Handler {
    r := chi.NewRouter()

    // Define user-related routes
    r.Get("/", users.HandleList(db))          // List all users
    r.Get("/{username}", users.HandleGetByName(db))   // Get a specific user by name
    r.Get("/id/{id}", users.HandleGetByID(db))   // Get a specific user by ID
    r.Post("/", users.HandleCreate(db))       // Create a new user
    r.Put("/{username}", users.HandleUpdate(db)) // Update a specific user by name
    r.Delete("/{username}", users.HandleDelete(db)) // Delete a specific user by name

    return r
}
