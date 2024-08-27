package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB(dbFileName string) error {
    var err error
    DB, err = sql.Open("sqlite", dbFileName)
    if err != nil {
        panic(err)
    }

    createTable := `
    CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        description TEXT,
        completed BOOLEAN,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );
    `

    if _, err := DB.Exec(createTable); err != nil {
        panic(err)
    }
    return nil
}

// CloseDB closes the database connection
func CloseDB() {
    if err := DB.Close(); err != nil {
        fmt.Printf("failed to close database: %v", err)
    }
}