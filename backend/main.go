package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

type Blog struct {
    Title string `json:"title"`
}

func main() {
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")

    dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ")/" + dbName
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/blogs", func(w http.ResponseWriter, r *http.Request) {
        rows, err := db.Query("SELECT title FROM blogs")
        if err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        defer rows.Close()

        var blogs []Blog
        for rows.Next() {
            var blog Blog
            rows.Scan(&blog.Title)
            blogs = append(blogs, blog)
        }

        json.NewEncoder(w).Encode(blogs)
    })

    log.Println("Backend running on :8080")
    http.ListenAndServe(":8080", nil)
}
