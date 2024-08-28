package cmd

import (
	"fmt"
	"go-task/db"
)

func UpdateTask(id int, title, description string) {
    _, err := db.DB.Exec("UPDATE tasks SET title = ?, description = ? WHERE id = ?", title, description, id)
    if err != nil {
        fmt.Println("Error updating task:", err)
    } else {
        fmt.Println("Task updated successfully!")
    }
}
