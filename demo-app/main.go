package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dbquery(w)
}

func dbquery(w http.ResponseWriter) {
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER_NAME")
	dbPass := os.Getenv("DB_USER_PASS")

	conStr := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":3306)/" + dbName
	log.Println(conStr)
	db, err := sql.Open("mysql", conStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Prepare statement for reading data
	rows, err := db.Query("SELECT id, name FROM movies")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var id int
	var name string

	html := "<!DOCTYPE html><html><head><meta charset=\"UTF-8\"><title>Demo App</title></head><body>"
	for rows.Next() {
		rows.Scan(&id, &name)
		html += fmt.Sprintf("<div>%d : %s </div>", id, name)
	}

	html += "</body></html>"
	fmt.Fprintf(w, "%s", html)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Start listenning at port 8080")
	http.ListenAndServe(":8080", nil)
}
