Task Tracker CLI
A command-line task management tool written in Go that allows you to track tasks with different statuses.

Features
- Add tasks with descriptions and unique IDs

- List tasks with optional status filtering

- Update task status (not-started, in-progress, completed)

- Delete tasks by ID

- Persistent storage using JSON file format

- Task metadata including creation and update timestamps

Installation
Clone the repository:

bash
git clone <repository-url>
cd Task-Tracker
Build the project:

bash
go build -o task-tracker main.go
Usage
Basic Commands
bash
# Add a new task
./task-tracker add <id> "<description>"

# List all tasks
./task-tracker list

# List tasks with specific status (not-started, in-progress, completed)
./task-tracker list not-started
./task-tracker list in-progress completed

# Update a task's status
./task-tracker update <id> <new-status>

# Delete a task
./task-tracker delete <id>
Examples
bash
# Add a task with ID 1
./task-tracker add 1 "Learn Go programming"

# Add another task
./task-tracker add 2 "Build a CLI application"

# List all tasks
./task-tracker list

# Update task 1 to in-progress
./task-tracker update 1 in-progress

# List only in-progress tasks
./task-tracker list in-progress

# Mark task 2 as completed
./task-tracker update 2 completed

# Delete task 1
./task-tracker delete 1
Task Status Options
The application supports three task statuses:

not-started - Task has been created but not started

in-progress - Task is currently being worked on

completed - Task has been finished

Data Storage
Tasks are automatically saved to a SavedTasks JSON file in the same directory. The file format preserves task information including:

Task ID

Description

Status

Creation timestamp

Last update timestamp

Project Structure
text
Task-Tracker/
    main.go                 # CLI entry point and command parsing
    handlers/
        command_handler.go  # Core task management functions
    tests/
        save_handler.go     # JSON save/load functionality
    task-data/              # Task data structures and validation
    SavedTasks              # Generated data file (after first use)
Dependencies
Standard Go libraries only

No external dependencies required

Development
The project is organized into separate packages:

handlers - Business logic for task operations

tests - Data persistence functions

task-data - Data structures and validation

Building from Source
bash
go mod tidy
go build -o task-tracker .