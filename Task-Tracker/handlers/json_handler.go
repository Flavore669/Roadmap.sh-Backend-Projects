package json_handler

//HACK: I should probably seperate this script into 2 scripts, 1 for save functionality. The other serves strictly for command functions that use the former script.
//TODO: Setup other methods (update)

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	taskConfig "github.com/Flavore669/Roadmap.sh-Backend-Projects/Task-Tracker/task-data"
)

type TaskJSON struct {
	Tasks []taskConfig.Task `json:"Tasks"`
}

var totalTasks TaskJSON
var tasksAsJSON []byte

func tasksContains(TargetID int) (int, error) {
	for index, task := range totalTasks.Tasks {
		if task.ID == TargetID {
			return index, nil
		}
	}

	return -1, errors.New("id doesn't exist")
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
		fmt.Printf("Task ID: %v, Description: %s, Progress: %s\n", task.ID, task.Description, task.TaskStatus)
	}
}

func DeleteTask(TargetID int) error {
	targetIndex, err := tasksContains(TargetID)

	if err != nil {
		return err
	}

	var tasksCopy TaskJSON
	for i := range totalTasks.Tasks {
		if i == targetIndex {
			continue
		}
		tasksCopy.Tasks = append(tasksCopy.Tasks, totalTasks.Tasks[i])
	}
	totalTasks = tasksCopy
	SaveData()

	return nil
}

func UpdateTask(targetID int, newStatus string) error {
	targetIndex, err := tasksContains(targetID)
	if err != nil {
		return err
	}

	_, err1 := taskConfig.IsValidStatus(newStatus)
	if err1 != nil {
		return err1
	}

	for i := range totalTasks.Tasks {
		if i == targetIndex {
			totalTasks.Tasks[i].TaskStatus = newStatus
		}
	}
	SaveData()

	return nil
}
