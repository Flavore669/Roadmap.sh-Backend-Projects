# Task Tracker CLI

A robust command-line task management tool written in Go that helps you track tasks with different statuses and persistent storage.

## Features

- **Add tasks** with descriptions and unique IDs
- **List tasks** with optional status filtering
- **Update task status** (not-started, in-progress, completed)
- **Delete tasks** by ID
- **Persistent storage** using JSON file format
- **Task metadata** including creation and update timestamps
- **Zero external dependencies** - uses standard Go libraries only

## Installation

### Prerequisites
- Go 1.16 or higher

### Build from Source
```bash
# Clone the repository
git clone <repository-url>
cd Task-Tracker

# Build the project
go build -o task-tracker main.go
```

### Usage
Basic Commands
```bash
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
```

Examples
```bash
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
```

### Task Status Options

The application supports three task statuses:

- not-started → Task has been created but not started

- in-progress → Task is currently being worked on

- completed → Task has been finished

### Data Storage

Tasks are automatically saved to a SavedTasks JSON file in the same directory.
The file format preserves task information including:

- Task ID

- Description

- Status

- Creation timestamp

- Last update timestamp
