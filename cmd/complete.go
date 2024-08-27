package cmd

import (
	"fmt"
	"task-manager-cli/db"
)

func CompleteTask(id int) {
    _, err := db.DB.Exec("UPDATE tasks SET completed = true WHERE id = ?", id)
    if err != nil {
        fmt.Println("Error completing task:", err)
    } else {
        fmt.Println("Task marked as completed!")
    }
}
