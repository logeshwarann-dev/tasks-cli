package services

import (
	"fmt"
	"time"

	"github.com/logeshwarann-dev/taskcli/internal/models"
	"github.com/logeshwarann-dev/taskcli/pkg/utils"
)

func AddTask(taskDesc string) error {

	var newTask models.Task
	var allTasks models.Tasks
	if utils.CheckIfFileExists(models.FilePath) {
		if err := utils.ReadJSONFile(models.FilePath, &allTasks); err != nil {
			return err
		}
	}

	lastId := getLastTaskId(allTasks)

	newTask.Id = lastId + 1
	newTask.Description = taskDesc
	newTask.CreatedAt = time.Now().Local()
	newTask.UpdatedAt = time.Now().Local()

	allTasks.TasksList = append(allTasks.TasksList, newTask)

	fmt.Println(allTasks.TasksList)

	if err := utils.WriteToFile(models.FilePath, allTasks); err != nil {
		return err
	}

	return nil

}

func getLastTaskId(tasksList models.Tasks) int {
	if len(tasksList.TasksList) == 0 {
		return 0
	}
	lastIndex := len(tasksList.TasksList) - 1
	lastTask := tasksList.TasksList[lastIndex]
	return lastTask.Id
}
