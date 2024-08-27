package cmd

import (
	"fmt"
	"task-manager-cli/db"
	"task-manager-cli/models"
)

func ListTasks() {
    rows, err := db.DB.Query("SELECT id, title, description, completed, created_at FROM tasks")
    if err != nil {
        fmt.Println("Error listing tasks:", err)
    }
    defer rows.Close()

    for rows.Next() {
        var task models.Task
        rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt)
        fmt.Printf("%d: %s - %s (Completed: %t)\n", task.ID, task.Title, task.Description, task.Completed)
    }
}
