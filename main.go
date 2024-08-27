package main

import (
	"os"
	"task-manager-cli/cmd"
	"task-manager-cli/db"

	"github.com/spf13/cobra"
)

func main() {
    db.InitDB()

    var rootCmd = &cobra.Command{Use: "task-manager"}

    var addCmd = &cobra.Command{
        Use:   "add",
        Short: "Add a new task",
        Run: func(cobraCmd *cobra.Command, args []string) {
            title, _ := cobraCmd.Flags().GetString("title")
            description, _ := cobraCmd.Flags().GetString("description")
            cmd.AddTask(title, description)
        },
    }

    addCmd.Flags().String("title", "", "Title of the task")
    addCmd.Flags().String("description", "", "Description of the task")
    rootCmd.AddCommand(addCmd)

    var listCmd = &cobra.Command{
        Use:   "list",
        Short: "List all tasks",
        Run: func(cobraCmd *cobra.Command, args []string) {
            cmd.ListTasks()
        },
    }
    rootCmd.AddCommand(listCmd)

    var updateCmd = &cobra.Command{
        Use:   "update",
        Short: "Update a task",
        Run: func(cobraCmd *cobra.Command, args []string) {
            id, _ := cobraCmd.Flags().GetInt("id")
            title, _ := cobraCmd.Flags().GetString("title")
            description, _ := cobraCmd.Flags().GetString("description")
            cmd.UpdateTask(id, title, description)
        },
    }
    updateCmd.Flags().Int("id", 0, "ID of the task")
    updateCmd.Flags().String("title", "", "New title of the task")
    updateCmd.Flags().String("description", "", "New description of the task")
    rootCmd.AddCommand(updateCmd)


    var deleteCmd = &cobra.Command{
        Use:   "delete",
        Short: "Delete a task",
        Run: func(cobraCmd *cobra.Command, args []string) {
            id, _ := cobraCmd.Flags().GetInt("id")
            cmd.DeleteTask(id)
        },
    }
    deleteCmd.Flags().Int("id", 0, "ID of the task")
    rootCmd.AddCommand(deleteCmd)

    var completeCmd = &cobra.Command{
        Use:   "complete",
        Short: "Mark a task as completed",
        Run: func(cobraCmd *cobra.Command, args []string) {
            id, _ := cobraCmd.Flags().GetInt("id")
            cmd.CompleteTask(id)
        },
    }
    completeCmd.Flags().Int("id", 0, "ID of the task")
    rootCmd.AddCommand(completeCmd)

    var searchCmd = &cobra.Command{
        Use:   "search",
        Short: "Search tasks by keyword",
        Run: func(cobraCmd *cobra.Command, args []string) {
            keyword, _ := cobraCmd.Flags().GetString("keyword")
            cmd.SearchTasks(keyword)
        },
    }
    searchCmd.Flags().String("keyword", "", "Keyword to search tasks")
    rootCmd.AddCommand(searchCmd)

    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
