package routes

import (
    "net/http"

    "github.com/CVWO/sample-go-app/internal/handlers/comments"
    "github.com/go-chi/chi/v5"
)

// CommentsRoutes sets up the routes for comment-related operations.
func CommentsRoutes() http.Handler {
    r := chi.NewRouter()

    // Define comment-related routes
    r.Get("/post/{postID}", comments.HandleList)  // List all comments by post ID
    r.Get("/{id}", comments.HandleGet)       // Get a specific comment by ID
    r.Post("/", comments.HandleCreate)       // Create a new comment
    r.Put("/{id}", comments.HandleUpdate)    // Update a specific comment by ID
    r.Delete("/{id}", comments.HandleDelete) // Delete a specific comment by ID

    return r
}
