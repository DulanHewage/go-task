package cmd

import (
	"go-task/db"
	"go-task/models"
	"testing"
)

func TestCompleteTask(t *testing.T) {
    db.InitDB("./test_tasks.db")
    defer db.CloseDB()
    defer db.DB.Exec("DROP TABLE IF EXISTS tasks")

    // Add a task to complete
    title := "Task to Complete"
    description := "This task will be completed"
    AddTask(title, description)

    // Complete the task
    var task models.Task
    err := db.DB.QueryRow("SELECT id FROM tasks WHERE title = ?", title).Scan(&task.ID)
    if err != nil {
        t.Fatalf("Failed to query task: %v", err)
    }
    CompleteTask(task.ID)

    // Verify the task is marked as completed
    err = db.DB.QueryRow("SELECT completed FROM tasks WHERE id = ?", task.ID).Scan(&task.Completed)
    if err != nil {
        t.Fatalf("Failed to query task: %v", err)
    }
    if !task.Completed.Bool {
        t.Errorf("Expected task to be completed")
    }
}