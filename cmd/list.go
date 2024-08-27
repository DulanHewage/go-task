package cmd

import (
	"fmt"
	"task-manager-cli/db"
	"task-manager-cli/models"
)

// ListTasks retrieves all tasks from the database and returns them.
func ListTasks() ([]models.Task, error) {
    rows, err := db.DB.Query("SELECT id, title, description, completed, created_at FROM tasks")
    if err != nil {
        return nil, fmt.Errorf("error listing tasks: %v", err)
    }
    defer rows.Close()

    var tasks []models.Task
    for rows.Next() {
        var task models.Task
        if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt); err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
        fmt.Printf("%d: %s - %s (Completed: %t)\n", task.ID, task.Title, task.Description, task.Completed)
    }
    return tasks, nil
}