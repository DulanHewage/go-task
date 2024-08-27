package cmd

import (
	"task-manager-cli/db"
	"task-manager-cli/models"
	"testing"
)

func TestDeleteTask(t *testing.T) {
    db.InitDB("./test_tasks.db")
    defer db.CloseDB()
    defer db.DB.Exec("DROP TABLE IF EXISTS tasks")

    // Add a task to delete
    title := "Task to Delete"
    description := "This task will be deleted"
    AddTask(title, description)

    // Delete the task
    var task models.Task
    err := db.DB.QueryRow("SELECT id FROM tasks WHERE title = ?", title).Scan(&task.ID)
    if err != nil {
        t.Fatalf("Failed to query task: %v", err)
    }
    DeleteTask(task.ID)

    // Verify the task is deleted
    err = db.DB.QueryRow("SELECT id FROM tasks WHERE id = ?", task.ID).Scan(&task.ID)
    if err == nil {
        t.Errorf("Expected task to be deleted")
    }
}