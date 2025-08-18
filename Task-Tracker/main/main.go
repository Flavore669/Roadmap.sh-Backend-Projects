package main

import (
	"fmt"
	"os"
)

type State int64

const (
	NotStarted State = iota
	InProgress
	Completed
)

//TODO: Create a seperate package for JSON data handling

type Task struct {
	ID    int
	Name  string
	State State
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Avaliable Commands: Add Task")
		fmt.Println("To Use Commands run - go run main.go {your_command} {additional_data}")
		os.Exit(0)
	}

	switch string(os.Args[1]) {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: No task description provided")
			fmt.Println("Usage: go run main.go add 'task description'")
			os.Exit(1)
		}
		fmt.Println("Added Task")
		err := addTask(os.Args[2])
		if err != nil {
			fmt.Printf("Error Adding Task, %s", os.Args[2])
		}

	default:
		fmt.Println("Please Use a Valid Command")
		fmt.Println("Avaliable Commands: Add Task")
		fmt.Println("To Use Commands run - go run main.go {your_command}")
	}
}

func addTask(task string) error {
	fmt.Print(task)
	return nil
}
