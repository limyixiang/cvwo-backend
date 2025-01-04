package database

import (
    "database/sql"
    "fmt"
    "os"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

type Database struct {
    *sql.DB
}

func GetDB() (*Database, error) {
    dsn := os.Getenv("JAWSDB_URL")
    if dsn == "" {
        return nil, fmt.Errorf("JAWSDB_URL environment variable is not set")
    }

    fmt.Println("Connecting to database with DSN:", dsn)

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, fmt.Errorf("error validating sql.Open arguments: %w", err)
    }

    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(5 * time.Minute)

    err = db.Ping()
    if err != nil {
        return nil, fmt.Errorf("error verifying connection to database with db.Ping: %w", err)
    }

    fmt.Println("Successfully connected to database!")
    return &Database{DB: db}, nil
}
