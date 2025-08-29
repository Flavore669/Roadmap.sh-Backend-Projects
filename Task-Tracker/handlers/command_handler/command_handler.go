package handlers

//HACK: I should probably seperate this script into 2 scripts, 1 for save functionality. The other serves strictly for command functions that use the former script.
//TODO: Implement listing tasks by status and that should be done! Then Refactor
import (
	"errors"
	"fmt"
	"strconv"
	"time"

	saveSystem "github.com/Flavore669/Roadmap.sh-Backend-Projects/Task-Tracker/handlers/save_handler"
	taskConfig "github.com/Flavore669/Roadmap.sh-Backend-Projects/Task-Tracker/task-data"
)

var tasksSaved taskConfig.TaskJSON

func init() {
	saveSystem.LoadData(&tasksSaved)
}

// Return desired format for the time
func return_time(target_time time.Time) string {
	return target_time.Month().String() + " " + strconv.Itoa(target_time.Day()) + " - " + strconv.Itoa(target_time.Year())
}

func tasksContains(TargetID int) (int, error) {
	for index, task := range tasksSaved.Tasks {
		if task.ID == TargetID {
			return index, nil
		}
	}

	return -1, errors.New("id doesn't exist")
}

func AddTask(taskID int, description string) error {
	var addedTask taskConfig.Task
	addedTask.ID = taskID
	addedTask.CreatedAt = time.Now()
	addedTask.UpdatedAt = time.Now()
	addedTask.TaskStatus = "not-started"
	addedTask.Description = description

	tasksSaved.Tasks = append(tasksSaved.Tasks, addedTask)

	err := saveSystem.SaveData(tasksSaved)
	if err != nil {
		return err
	}

	return nil
}

func ListSavedTasks(flags ...string) error {
	tasks, err := saveSystem.LoadData(&tasksSaved)
	if err != nil {
		return err
	}

	for _, task := range tasks {
		valid := false
		if len(flags) > 0 { // If there are any flags
			for f := range flags { // Check if this task's status matches any element in flags
				err := taskConfig.IsValidStatus(flags[f])
				if err != nil {
					return err
				}

				if flags[f] == task.TaskStatus {
					valid = true
				}
			}
		} else {
			valid = true
		}

		if !valid {
			continue
		}

		fmt.Printf("Task ID: %v, Description: %s, Status: %s, Time Created: %s, Time Updated %s\n",
			task.ID, task.Description, task.TaskStatus, return_time(task.CreatedAt), return_time(task.UpdatedAt))
	}

	return nil
}

func DeleteTask(TargetID int) error {
	targetIndex, err := tasksContains(TargetID)

	if err != nil {
		return err
	}

	// Create a copy of the tasksSaved that doesn't include the deleted task
	var tasksCopy taskConfig.TaskJSON
	for i := range tasksSaved.Tasks {
		if i == targetIndex {
			continue
		}
		tasksCopy.Tasks = append(tasksCopy.Tasks, tasksSaved.Tasks[i])
	}
	tasksSaved = tasksCopy
	saveSystem.SaveData(tasksSaved)

	return nil
}

func UpdateTask(targetID int, newStatus string) error {
	targetIndex, err := tasksContains(targetID) // Find Index
	if err != nil {
		return err
	}

	err1 := taskConfig.IsValidStatus(newStatus)
	if err1 != nil {
		return err1
	}

	for i := range tasksSaved.Tasks {
		if i == targetIndex {
			tasksSaved.Tasks[i].TaskStatus = newStatus
			tasksSaved.Tasks[i].UpdatedAt = time.Now()
		}
	}
	saveSystem.SaveData(tasksSaved)

	return nil
}
