package cmd

import (
	"task-manager-cli/db"
	"task-manager-cli/models"
	"testing"
)

func TestUpdateTask(t *testing.T) {
    db.InitDB("./test_tasks.db")
    defer db.CloseDB()

    // Add a task to update
    title := "Task to Update"
    description := "This task will be updated"
    AddTask(title, description)

    // Update the task
    var task models.Task
    err := db.DB.QueryRow("SELECT id FROM tasks WHERE title = ?", title).Scan(&task.ID)
    if err != nil {
        t.Fatalf("Failed to query task: %v", err)
    }
    newTitle := "Updated Task"
    newDescription := "This task has been updated"
    UpdateTask(task.ID, newTitle, newDescription)

    // Verify the task is updated
    err = db.DB.QueryRow("SELECT title, description FROM tasks WHERE id = ?", task.ID).Scan(&task.Title, &task.Description)
    if err != nil {
        t.Fatalf("Failed to query task: %v", err)
    }
    if task.Title != newTitle {
        t.Errorf("Expected title %s, got %s", newTitle, task.Title)
    }
    if task.Description != newDescription {
        t.Errorf("Expected description %s, got %s", newDescription, task.Description)
    }
}