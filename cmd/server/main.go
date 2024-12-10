package main

import (
	// "fmt"
	"log"

	"github.com/CVWO/sample-go-app/internal/database"
	// _ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.GetDB()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()	
}
