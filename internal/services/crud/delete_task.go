package services

import (
	"errors"
	"fmt"
	"slices"

	"github.com/logeshwarann-dev/taskcli/internal/models"
	"github.com/logeshwarann-dev/taskcli/pkg/utils"
)

func DeleteTask(taskId int, filePath string) error {

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
			allTasks.TasksList = slices.Delete(allTasks.TasksList, index, index+1)
			isTaskIdPresent = true
		}
	}
	// fmt.Println("After Deleting: ", allTasks.TasksList)
	if !isTaskIdPresent {
		return errors.New("task id not found")
	}
	if err := utils.WriteToFile(filePath, allTasks); err != nil {
		return err
	}

	return nil

}

func PrintDeletedTask(taskId int) {
	fmt.Printf("Task Deleted Successfully (ID: %d)\n", taskId)
}
