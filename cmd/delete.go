package cmd

import (
	"fmt"
	"go-task/db"
)

func DeleteTask(id int) {
    _, err := db.DB.Exec("DELETE FROM tasks WHERE id = ?", id)
    if err != nil {
        fmt.Println("Error deleting task:", err)
    } else {
        fmt.Println("Task deleted successfully!")
    }
}
