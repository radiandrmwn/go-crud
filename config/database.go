package config

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
    dsn := "root:root@/go_products?parseTime=true"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }

    DB = db
    log.Println("Database connected")
}
