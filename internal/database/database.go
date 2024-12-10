package database

import (
    "database/sql"
    "fmt"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

type Database struct {
    *sql.DB
}

func GetDB() (*Database, error) {
    password := os.Getenv("MYSQL_PASSWORD")
    if password == "" {
        return nil, fmt.Errorf("MYSQL_PASSWORD environment variable is not set")
    }

    db, err := sql.Open("mysql", "root:"+password+"@tcp(localhost:3306)/testdb")
    if err != nil {
        return nil, fmt.Errorf("error validating sql.Open arguments: %w", err)
    }

    err = db.Ping()
    if err != nil {
        return nil, fmt.Errorf("error verifying connection to database with db.Ping: %w", err)
    }

    fmt.Println("Successfully connected to database!")
    return &Database{DB: db}, nil
}
