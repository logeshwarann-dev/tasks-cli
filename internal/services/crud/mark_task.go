package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/logeshwarann-dev/taskcli/internal/models"
	"github.com/logeshwarann-dev/taskcli/pkg/utils"
)

func MarkStatusOfTask(taskId int, status string, filePath string) error {

	var allTasks models.Tasks
	if utils.CheckIfFileExists(filePath) {
		isEmpty, err := utils.IsFileEmpty(filePath)
		if err != nil {
			return err
		}
		if !isEmpty { // Read file only if its not empty
			if err := utils.ReadJSONFile(filePath, &allTasks); err != nil {
				return err
			}
		}
	}

	var isTaskIdPresent bool
	for index, eachTask := range allTasks.TasksList {
		if eachTask.Id == taskId {
			eachTask.Status = status
			eachTask.UpdatedAt = time.Now().Local()
			allTasks.TasksList[index] = eachTask
			fmt.Println("Updated Task: ", allTasks.TasksList[index])
			isTaskIdPresent = true
		}
	}
	if !isTaskIdPresent {
		return errors.New("task id not found")
	}
	if err := utils.WriteToFile(filePath, allTasks); err != nil {
		return err
	}

	return nil
}

func PrintTaskStatus(taskId int) {
	fmt.Printf("Task Updated Successfully (ID: %d)\n", taskId)
}
