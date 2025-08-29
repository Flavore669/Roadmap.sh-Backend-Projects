package tests

import (
	"encoding/json"
	"fmt"
	"os"

	taskConfig "github.com/Flavore669/Roadmap.sh-Backend-Projects/Task-Tracker/task-data"
)

// Save inputted tasksSaved to a JSON file
func SaveData(tasksSaved taskConfig.TaskJSON) error {
	// Marshal JSON
	tasksAsJSON, err := json.MarshalIndent(tasksSaved, "", "\t")
	if err != nil {
		fmt.Printf("Error is %s", err)
		return err
	}

	// Write to File
	os.WriteFile("SavedTasks", tasksAsJSON, 0666)
	return nil
}

// Load the json file into the reference for tasksSaved
func LoadData(tasksSaved *taskConfig.TaskJSON) ([]taskConfig.Task, error) {
	// Read the file
	tasksAsJSON, err := os.ReadFile("SavedTasks")
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return nil, err
	}

	// Unmarhsal JSON
	err = json.Unmarshal(tasksAsJSON, &tasksSaved)
	if err != nil {
		fmt.Printf("Error is %s", err)
	}

	return tasksSaved.Tasks, nil
}
