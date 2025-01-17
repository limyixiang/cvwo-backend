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
    r.Get("/post/{postID}", comments.HandleList)                                // List all comments by post ID
    r.Get("/{id}", comments.HandleGet)                                          // Get a specific comment by ID
    r.Post("/", comments.HandleCreate)                                          // Create a new comment
    r.Patch("/{id}", comments.HandleUpdate)                                     // Update a specific comment by ID
    r.Patch("/{id}/like", comments.HandleLike)                                  // Like a specific comment by ID
    r.Patch("/{id}/unlike", comments.HandleUnlike)                              // Dislike a specific comment by ID
    r.Patch("/{id}/dislike", comments.HandleDislike)                            // Dislike a specific comment by ID
    r.Patch("/{id}/undislike", comments.HandleUndislike)                        // Undislike a specific comment by ID
    r.Delete("/{id}", comments.HandleDelete)                                    // Delete a specific comment by ID
    r.Get("/{id}/checklike/{userID}", comments.HandleCheckLikedByUser)          // Check if a user has liked a specific comment by ID
    r.Get("/{id}/checkdislike/{userID}", comments.HandleCheckDislikedByUser)    // Check if a user has disliked a specific comment by ID

    return r
}
