package services

import (
	"fmt"

	"github.com/logeshwarann-dev/taskcli/internal/models"
	"github.com/logeshwarann-dev/taskcli/pkg/utils"
)

func ListTask(status string, filePath string) error {

	var allTasks models.Tasks
	var outputTasks []models.Task
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

	if status == models.EmptyStatus {
		fmt.Println(allTasks.TasksList)
		return nil
	}

	for _, eachTask := range allTasks.TasksList {
		if eachTask.Status == status {
			outputTasks = append(outputTasks, eachTask)
		}
	}
	fmt.Println(outputTasks)
	return nil

}

func GetStatusFlag(args []string) string {
	if len(args) == 2 {
		return args[1]
	}
	return ""
}
