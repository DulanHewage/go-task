package cmd

import (
	"task-manager-cli/db"
	"testing"
)

func TestListTasks(t *testing.T) {
    db.InitDB("./test_tasks.db")
    defer db.CloseDB()
    defer db.DB.Exec("DROP TABLE IF EXISTS tasks")

    // Add a task to list
    title := "Task to List"
    description := "This task will be listed"
    AddTask(title, description)

    // List tasks
    tasks, err := ListTasks()
    if err != nil {
        t.Fatalf("Failed to list tasks: %v", err)
    }

    // Verify the task is listed
    found := false
    for _, task := range tasks {
        if task.Title == title && task.Description == description {
            found = true
            break
        }
    }
    if !found {
        t.Errorf("Expected task to be listed")
    }
}