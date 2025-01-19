package routes

import (
	"net/http"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/handlers/threads"
	"github.com/go-chi/chi/v5"
)

// ThreadsRoutes sets up the routes for thread-related operations.
func ThreadsRoutes(db *database.Database) http.Handler {
    r := chi.NewRouter()

    // Define thread-related routes
    r.Get("/", threads.HandleList(db))          // List all threads
    r.Get("/{id}", threads.HandleGet(db))       // Get a specific thread by ID
    r.Post("/", threads.HandleCreate(db))       // Create a new thread
    r.Put("/{id}", threads.HandleUpdate(db))    // Update a specific thread by ID
    r.Delete("/{id}", threads.HandleDelete(db)) // Delete a specific thread by ID

    return r
}
