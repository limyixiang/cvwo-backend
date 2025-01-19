package routes

import (
	"net/http"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/handlers/posts"
	"github.com/go-chi/chi/v5"
)

// PostsRoutes sets up the routes for post-related operations.
func PostsRoutes(db *database.Database) http.Handler {
    r := chi.NewRouter()

    // Define post-related routes
    r.Get("/", posts.HandleList(db))                                            // List all posts
    r.Get("/category/{id}", posts.HandleListByCategory(db))                     // List all posts in a specific category
    r.Get("/{id}", posts.HandleGet(db))                                         // Get a specific post by ID
    r.Patch("/{id}/updatetime", posts.HandleUpdateLastUpdated(db))              // Update the time of a specific post by ID
    r.Post("/", posts.HandleCreate(db))                                         // Create a new post
    r.Patch("/{id}", posts.HandleUpdate(db))                                    // Update a specific post by ID
    r.Patch("/{id}/like", posts.HandleLike(db))                                 // Like a specific post by ID
    r.Patch("/{id}/unlike", posts.HandleUnlike(db))                             // Dislike a specific post by ID
    r.Patch("/{id}/dislike", posts.HandleDislike(db))                           // Dislike a specific post by ID
    r.Patch("/{id}/undislike", posts.HandleUndislike(db))                       // Undislike a specific post by ID
    r.Delete("/{id}", posts.HandleDelete(db))                                   // Delete a specific post by ID
    r.Get("/{id}/checklike/{userID}", posts.HandleCheckLikedByUser(db))         // Check if a user has liked a specific post by ID
    r.Get("/{id}/checkdislike/{userID}", posts.HandleCheckDislikedByUser(db))   // Check if a user has disliked a specific post by ID

    return r
}
