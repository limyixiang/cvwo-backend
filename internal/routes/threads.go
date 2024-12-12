package routes

import (
    "net/http"

    "github.com/CVWO/sample-go-app/internal/handlers/threads"
    "github.com/go-chi/chi/v5"
)

// ThreadsRoutes sets up the routes for thread-related operations.
func ThreadsRoutes() http.Handler {
    r := chi.NewRouter()

    // Define thread-related routes
    r.Get("/", threads.HandleList)          // List all threads
    r.Get("/{id}", threads.HandleGet)       // Get a specific thread by ID
    r.Post("/", threads.HandleCreate)       // Create a new thread
    r.Put("/{id}", threads.HandleUpdate)    // Update a specific thread by ID
    r.Delete("/{id}", threads.HandleDelete) // Delete a specific thread by ID

    return r
}
