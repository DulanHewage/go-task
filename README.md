# Task Manager CLI with GoLang and SQLite 📟

The Task Manager CLI is a command-line interface (CLI) tool built with GoLang. Users can add, view, update, and delete tasks, all of which are stored in a SQLite database.

## Features

- Add Task: Add a new task with a title and description.
- List Tasks: View all tasks with details.
- Update Task: Modify existing tasks.
- Delete Task: Remove a task from the list.
- Mark as Done: Mark a task as completed.
- Search Tasks: Search tasks by keyword.

## Directory Structure

- `cmd/`: Contains Go files for different commands (add, list, update, etc.).
- `db/`: Manages database connection and migrations.
- `models/`: Defines the Task model and related database operations.
- `main.go`: The entry point of the application.

## Getting Started

### Prerequisites

- Go 1.21.2 or later
- SQLite

### Installation

1. **Clone the repository**:

   ```sh
   git clone https://github.com/dulanhewage/go-task.git
   cd go-task
   ```

2. **Install dependencies**:

   ```sh
   go mod tidy
   ```

3. **Build the project**:

   ```sh
   go build -o go-task
   ```

### Usage

After building the project, you can use the CLI tool by running the generated binary:

```sh
./go-task --help
```

### Example

Here is an example of how to use the Task Manager CLI:

1. Add a new task:

```sh
./go-task add --title "Buy groceries" --description "Milk, Bread, Eggs"
```

2. List all tasks:

```sh
./go-task list
```

3. Update a task:

```sh
./go-task update --id 1 --title "Buy groceries and fruits" --description "Milk, Bread, Eggs, Apples"
```

4. Mark a task as completed:

```sh
./go-task complete --id 1
```

5. Search tasks by keyword:

```sh
./go-task search --keyword "groceries"
```

6. Delete a task:

```sh
./go-task delete --id 1
```
