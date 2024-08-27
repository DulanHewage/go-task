package cmd

import (
	"fmt"
	"task-manager-cli/db"
	"task-manager-cli/models"
)

func SearchTasks(keyword string) {
    rows, err := db.DB.Query("SELECT id, title, description FROM tasks WHERE title LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
    if err != nil {
        fmt.Println("Error searching tasks:", err)
    }
    defer rows.Close()

    for rows.Next() {
        var task models.Task
        rows.Scan(&task.ID, &task.Title, &task.Description)
        fmt.Printf("%d: %s - %s\n", task.ID, task.Title, task.Description)
    }
}
