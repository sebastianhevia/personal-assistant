package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

var Database *sql.DB

func InitDB(filepath string) *sql.DB {
    var err error
    Database, err = sql.Open("sqlite3", filepath)
    if err != nil {
        log.Fatal(err)
    }

    if err := Database.Ping(); err != nil {
        log.Fatal(err)
    }

    return Database
}

