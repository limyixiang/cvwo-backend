package main

import (
    "log"
    "net/http"
    "os"

    "github.com/CVWO/sample-go-app/internal/database"
    "github.com/CVWO/sample-go-app/internal/router"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/cors"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Initialize the database connection
    db, err := database.GetDB()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    // Set up the router
    r := chi.NewRouter()

    // Enable CORS
    r.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           300, // Maximum value not ignored by any of major browsers
    }))

    // Set up routes
    r.Mount("/", router.Setup())

    // Get the port from the environment variables or use a default port
    port := os.Getenv("MYSQL_PORT")
    if port == "" {
        port = os.Getenv("PORT")
    }

    if port == "" {
        port = "8080"
    }

    // Start the HTTP server
    log.Printf("Starting server on port %s...", port)
    if err := http.ListenAndServe(":"+port, r); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
