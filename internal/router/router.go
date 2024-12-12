package router

import (
	"net/http"

	"github.com/CVWO/sample-go-app/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Setup() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	setUpRoutes(r)
	return r
}

func setUpRoutes(r chi.Router) {
	r.Route("/api", func(r chi.Router) {
        r.Mount("/users", routes.UsersRoutes())
        r.Mount("/posts", routes.PostsRoutes())
        r.Mount("/threads", routes.ThreadsRoutes())
        r.Mount("/comments", routes.CommentsRoutes())
        r.Mount("/categories", routes.CategoriesRoutes())
    })

    // Add a health check endpoint
    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    })
}
