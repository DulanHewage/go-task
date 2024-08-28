package cmd

import (
	"fmt"
	"go-task/db"
)

func AddTask(title, description string) {
    _, err := db.DB.Exec("INSERT INTO tasks (title, description) VALUES (?, ?)", title, description)
    if err != nil {
        fmt.Println("Error adding task:", err)
    } else {
        fmt.Println("Task added successfully!")
    }
}
