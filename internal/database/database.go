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
    // password := os.Getenv("MYSQL_PASSWORD")
	// fmt.Println(password)
    // if password == "" {
    //     return nil, fmt.Errorf("MYSQL_PASSWORD environment variable is not set")
    // }

    // db, err := sql.Open("mysql", "root:"+password+"@tcp(localhost:3306)/testdb")
    // if err != nil {
    //     return nil, fmt.Errorf("error validating sql.Open arguments: %w", err)
    // }

    // mysql://fdq8lf5t43k2trxc:ur0td788zggb5u3g@l3855uft9zao23e2.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306/f0keggdl661acodt

    dsn := os.Getenv("JAWSDB_URL")
    if dsn == "" {
        return nil, fmt.Errorf("JAWSDB_URL environment variable is not set")
    }

    fmt.Println("Connecting to database with DSN:", dsn)

    db, err := sql.Open("mysql", dsn)
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
