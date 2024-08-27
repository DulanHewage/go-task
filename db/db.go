package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite", "./tasks.db")
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
}
