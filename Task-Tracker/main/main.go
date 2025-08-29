package main

//TODO: Rename the commandHandler to commandHandler when using new saver
import (
	"fmt"
	"os"
	"strconv"

	commandHandler "github.com/Flavore669/Roadmap.sh-Backend-Projects/Task-Tracker/handlers/command_handler"
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
	// If user doesn't enter enough arguements, exit
	if len(os.Args) < 2 {
		fmt.Println("Avaliable Commands: Add Task")
		fmt.Println("To Use Commands run - go run main.go {your_command} {id} {additional_data}")
		os.Exit(0)
	}

	/*
		All cases follow the same formula.
		Check if there are a correct number of arguements.
		Process arguements using commandHandler -> Return any errors that may occur
	*/
	switch string(os.Args[1]) {
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Error: No ID or description provided")
			fmt.Println("Correct Usage: go run main.go add 'id' 'task description'")
			os.Exit(1)
		}

		ID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Error: Invalid ID '%s'\n", os.Args[2])
			os.Exit(1)
		}

		err = commandHandler.AddTask(ID, os.Args[3])
		if err != nil {
			fmt.Printf("Error adding task: %s\n", err)
			os.Exit(1)
		}
		fmt.Println("Added task successfully")

	case "list":
		additionalArgs := os.Args[2:]
		err2 := commandHandler.ListSavedTasks(additionalArgs...)
		if err2 != nil {
			fmt.Printf("Error Listing Tasks: %s", err2)
		}

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: No ID or Status Provided")
			fmt.Println("Correct Usage: go run main.go delete 'id'")
			os.Exit(1)
		}

		ID, err := strconv.Atoi(os.Args[2]) // Convert ID arguement to an int
		if err != nil {
			fmt.Printf("Error Deleting Task: %s", err)
			os.Exit(1)
		}

		err3 := commandHandler.DeleteTask(ID)
		if err3 != nil {
			fmt.Printf("Error: %s", err3)
			os.Exit(1)
		}
		fmt.Println("Deleted Task")
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Error: No ID Provided")
			fmt.Println("Correct Usage: go run main.go update 'id' 'New Status'")
			os.Exit(1)
		}

		ID, err := strconv.Atoi(os.Args[2]) // Convert ID arguement to an int
		if err != nil {
			fmt.Printf("Error Deleting Task: %s", err)
			os.Exit(1)
		}

		err3 := commandHandler.UpdateTask(ID, os.Args[3])
		if err3 != nil {
			fmt.Printf("Error: %s", err3)
			os.Exit(1)
		}

		fmt.Println("Updated Task")

	default:
		fmt.Println("Please Use a Valid Command")
		fmt.Println("Avaliable Commands: Add Task")
		fmt.Println("To Use Commands run - go run main.go {your_command}")
	}
}
