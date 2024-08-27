# Task Manager CLI with GoLang and SQLite

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
   git clone https://github.com/dulanhewage/task-manager-cli.git
   cd task-manager-cli
   ```

2. **Install dependencies**:

   ```sh
   go mod tidy
   ```

3. **Build the project**:

   ```sh
   go build -o task-manager-cli
   ```

### Usage

After building the project, you can use the CLI tool by running the generated binary:

```sh
./task-manager-cli --help
```

### Example

Here is an example of how to use the Task Manager CLI:

1. Add a new task:

```sh
./task-manager-cli add --title "Buy groceries" --description "Milk, Bread, Eggs"
```

2. List all tasks:

```sh
./task-manager-cli list
```

3. Update a task:

```sh
./task-manager-cli update --id 1 --title "Buy groceries and fruits" --description "Milk, Bread, Eggs, Apples"
```

4. Mark a task as completed:

```sh
./task-manager-cli complete --id 1
```

5. Search tasks by keyword:

```sh
./task-manager-cli search --keyword "groceries"
```

6. Delete a task:

```sh
./task-manager-cli delete --id 1
```
