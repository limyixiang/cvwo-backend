package main

import (
	"database/sql"
	"fmt"
	// "log"
	// "net/http"
	"os"

	// "github.com/CVWO/sample-go-app/internal/router"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// r := router.Setup()
	// fmt.Print("Listening on port 8000 at http://localhost:8000!")

	// log.Fatalln(http.ListenAndServe(":8000", r))

	fmt.Println(os.Getenv("MYSQL_PASSWORD"))
	password := os.Getenv("MYSQL_PASSWORD")
	db, err := sql.Open("mysql", "root:"+password+"@tcp(localhost:3306)/testdb")

	if err != nil {
		fmt.Println("Error validating sql.Open arguments")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error verifying connection to database with db.Ping")
		panic(err.Error())
	}

	fmt.Println("Successfully connected to database!")
}
