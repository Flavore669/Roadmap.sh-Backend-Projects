package tests

import (
	"encoding/json"
	"fmt"
	"os"

	taskConfig "github.com/Flavore669/Roadmap.sh-Backend-Projects/Task-Tracker/task-data"
)

var tasksAsJSON []byte

func SaveData(tasksSaved taskConfig.TaskJSON) {
	var err error
	tasksAsJSON, err = json.MarshalIndent(tasksSaved, "", "\t")
	if err != nil {
		fmt.Printf("Error is %s", err)
	}

	os.WriteFile("SavedTasks", tasksAsJSON, 0666)
}

func LoadData(tasksSaved *taskConfig.TaskJSON) []taskConfig.Task {
	var err error

	// Read the file
	tasksAsJSON, err = os.ReadFile("SavedTasks")
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return nil
	}

	// Unmarhsal JSON
	err = json.Unmarshal(tasksAsJSON, &tasksSaved)
	if err != nil {
		fmt.Printf("Error is %s", err)
	}

	return tasksSaved.Tasks
}
