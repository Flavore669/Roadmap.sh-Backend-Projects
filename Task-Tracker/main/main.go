package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	taskFile "github.com/Flavore669/Roadmap.sh-Backend-Projects/Task-Tracker/handlers"
	taskConfig "github.com/Flavore669/Roadmap.sh-Backend-Projects/Task-Tracker/task-data"
)

type TaskState int64

const (
	NotStarted TaskState = iota
	InProgress
	Completed
)

type Task struct {
	ID        int
	Name      string
	TaskState TaskState
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Avaliable Commands: Add Task")
		fmt.Println("To Use Commands run - go run main.go {your_command} {id} {additional_data}")
		os.Exit(0)
	}

	taskFile.LoadData()

	switch string(os.Args[1]) {
	case "add":
		//HACK: Refactor so it has similar structure as other cases
		if len(os.Args) < 4 {
			fmt.Println("Error: No ID or description provided")
			fmt.Println("Correct Usage: go run main.go add 'id' 'task description'")
			os.Exit(1)
		}

		fmt.Println("Added Task")

		ID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Error Adding Task: %s", os.Args[2])
		}

		err1 := addTask(ID, os.Args[3])
		if err1 != nil {
			fmt.Printf("Error Adding Task, %s", os.Args[3])
		}

	case "list":
		additionalArgs := os.Args[2:]

		err2 := taskFile.ListSavedTasks(additionalArgs...)
		if err2 != nil {
			fmt.Printf("Error Listing Tasks: %s", err2)
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: No ID or Status Provided")
			fmt.Println("Correct Usage: go run main.go delete 'id'")
			os.Exit(1)
		}
		ID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Error Deleting Task: %s", err)
		}
		err3 := taskFile.DeleteTask(ID)
		if err3 != nil {
			fmt.Printf("Error: %s", err3)
		}
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Error: No ID Provided")
			fmt.Println("Correct Usage: go run main.go update 'id' 'New Status'")
			os.Exit(1)
		}

		ID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Error Deleting Task: %s", err)
		}

		err3 := taskFile.UpdateTask(ID, os.Args[3])
		if err3 != nil {
			fmt.Printf("Error: %s", err3)
		}

	default:
		fmt.Println("Please Use a Valid Command")
		fmt.Println("Avaliable Commands: Add Task")
		fmt.Println("To Use Commands run - go run main.go {your_command}")
	}
}

func addTask(taskID int, description string) error {
	var addedTask taskConfig.Task
	addedTask.ID = taskID
	addedTask.CreatedAt = time.Now()
	addedTask.UpdatedAt = time.Now()
	addedTask.TaskStatus = "not-started"
	addedTask.Description = description

	taskFile.AddTask(addedTask)
	return nil
}
