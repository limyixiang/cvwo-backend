package router

import (
	"net/http"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Setup(db *database.Database) chi.Router {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    setUpRoutes(r, db)
    return r
}

func setUpRoutes(r chi.Router, db *database.Database) {
    r.Route("/api", func(r chi.Router) {
        r.Mount("/users", routes.UsersRoutes(db))
        r.Mount("/posts", routes.PostsRoutes(db))
        r.Mount("/threads", routes.ThreadsRoutes(db))
        r.Mount("/comments", routes.CommentsRoutes(db))
        r.Mount("/categories", routes.CategoriesRoutes(db))
    })

    // Add a health check endpoint
    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    })
}
