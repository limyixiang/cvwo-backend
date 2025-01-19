package routes

import (
    "net/http"

    "github.com/CVWO/sample-go-app/internal/database"
    "github.com/CVWO/sample-go-app/internal/handlers/comments"
    "github.com/go-chi/chi/v5"
)

// CommentsRoutes sets up the routes for comment-related operations.
func CommentsRoutes(db *database.Database) http.Handler {
    r := chi.NewRouter()

    // Define comment-related routes
    r.Get("/post/{postID}", comments.HandleList(db))                                // List all comments by post ID
    r.Get("/{id}", comments.HandleGet(db))                                          // Get a specific comment by ID
    r.Post("/", comments.HandleCreate(db))                                          // Create a new comment
    r.Patch("/{id}", comments.HandleUpdate(db))                                     // Update a specific comment by ID
    r.Patch("/{id}/like", comments.HandleLike(db))                                  // Like a specific comment by ID
    r.Patch("/{id}/unlike", comments.HandleUnlike(db))                              // Dislike a specific comment by ID
    r.Patch("/{id}/dislike", comments.HandleDislike(db))                            // Dislike a specific comment by ID
    r.Patch("/{id}/undislike", comments.HandleUndislike(db))                        // Undislike a specific comment by ID
    r.Delete("/{id}", comments.HandleDelete(db))                                    // Delete a specific comment by ID
    r.Get("/{id}/checklike/{userID}", comments.HandleCheckLikedByUser(db))          // Check if a user has liked a specific comment by ID
    r.Get("/{id}/checkdislike/{userID}", comments.HandleCheckDislikedByUser(db))    // Check if a user has disliked a specific comment by ID

    return r
}
