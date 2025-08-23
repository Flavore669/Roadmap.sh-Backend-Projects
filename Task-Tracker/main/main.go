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
		if len(os.Args) < 4 {
			fmt.Println("Error: No task provided")
			fmt.Println("Usage: go run main.go add 'id' 'task description'")
			os.Exit(1)
		}

		fmt.Println("Added Task")

		ID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Error Adding Task, %s", os.Args[2])
		}

		//TODO Add support for a description arguement in add task
		err1 := addTask(ID, os.Args[3])
		if err1 != nil {
			fmt.Printf("Error Adding Task, %s", os.Args[3])
		}

	case "list":
		taskFile.ListSavedTasks()
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
	addedTask.CreatedAt = time.Now()
	var taskStatus taskConfig.TaskStatus = taskConfig.TaskStatus(NotStarted)
	addedTask.TaskStatus = taskConfig.TaskStatusToString[taskStatus]
	addedTask.Description = description

	taskFile.AddTask(addedTask)
	return nil
}
