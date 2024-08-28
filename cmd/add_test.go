package cmd

import (
	"go-task/db"
	"go-task/models"
	"testing"
)

func TestAddTask(t *testing.T) {
    // Initialize the database (use a test database)
    db.InitDB("./test_tasks.db")
    defer db.CloseDB()
    defer db.DB.Exec("DROP TABLE IF EXISTS tasks")

    // Add a task
    title := "Test Task"
    description := "This is a test task"
    AddTask(title, description)

    // Query the database to check if the task was added
    var task models.Task
    err := db.DB.QueryRow("SELECT id, title, description FROM tasks WHERE title = ?", title).Scan(&task.ID, &task.Title, &task.Description)
    if err != nil {
        t.Fatalf("Failed to query task: %v", err)
    }

    // Check if the task details are correct
    if task.Title != title {
        t.Errorf("Expected title %s, got %s", title, task.Title)
    }
    if task.Description != description {
        t.Errorf("Expected description %s, got %s", description, task.Description)
    }
}