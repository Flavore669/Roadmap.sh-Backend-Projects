package json_handler

//TODO: Setup other identifiers for tasks (description). Setup other methods (list, delete)

import (
	"encoding/json"
	"fmt"
	"os"

	taskConfig "github.com/Flavore669/Roadmap.sh-Backend-Projects/Task-Tracker/task-data"
)

type TaskJSON struct {
	Tasks []taskConfig.Task `json:"Tasks"`
}

var totalTasks TaskJSON
var tasksAsJSON []byte

func contains(otherTask taskConfig.Task) bool {
	for _, task := range totalTasks.Tasks {
		if task.ID == otherTask.ID {
			return true
		}
	}

	return false
}

func AddTask(task taskConfig.Task) {
	totalTasks.Tasks = append(totalTasks.Tasks, task)
	SaveData()
}

func SaveData() {
	var err error
	tasksAsJSON, err = json.Marshal(totalTasks)
	if err != nil {
		fmt.Printf("Error is %s", err)
	}

	os.WriteFile("SavedTasks", tasksAsJSON, 0666)
}

func LoadData() []taskConfig.Task {
	var err error

	// Read the file
	tasksAsJSON, err = os.ReadFile("SavedTasks")
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return nil
	}

	// Unmarhsal JSON
	err = json.Unmarshal(tasksAsJSON, &totalTasks)
	if err != nil {
		fmt.Printf("Error is %s", err)
	}

	return totalTasks.Tasks
}

func ListSavedTasks() {
	var tasks []taskConfig.Task = LoadData()

	for _, task := range tasks {
		fmt.Printf("Task: %s\n", task.Description)
	}
}
