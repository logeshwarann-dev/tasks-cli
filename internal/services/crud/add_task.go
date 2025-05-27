package services

import (
	"fmt"
	"time"

	"github.com/logeshwarann-dev/taskcli/internal/models"
	"github.com/logeshwarann-dev/taskcli/pkg/utils"
)

func AddTask(taskDesc string, filePath string) (int, error) {

	var newTask models.Task
	var allTasks models.Tasks
	if utils.CheckIfFileExists(filePath) {
		isEmpty, err := utils.IsFileEmpty(filePath)
		if err != nil {
			return 0, err
		}
		if !isEmpty { // Read file only if its not empty
			if err := utils.ReadJSONFile(filePath, &allTasks); err != nil {
				return 0, err
			}
		}
	}

	lastId := getLastTaskId(allTasks)

	newTask.Id = lastId + 1
	newTask.Description = taskDesc
	newTask.Status = models.StatusTodo
	newTask.CreatedAt = time.Now().Local()
	newTask.UpdatedAt = time.Now().Local()

	allTasks.TasksList = append(allTasks.TasksList, newTask)

	fmt.Println(allTasks.TasksList)

	if err := utils.WriteToFile(filePath, allTasks); err != nil {
		return 0, err
	}

	return newTask.Id, nil

}

func getLastTaskId(tasksList models.Tasks) int {
	if len(tasksList.TasksList) == 0 {
		return 0
	}
	lastIndex := len(tasksList.TasksList) - 1
	lastTask := tasksList.TasksList[lastIndex]
	return lastTask.Id
}

func PrintAddedTask(taskId int) {
	fmt.Printf("Task Added Successfully (ID: %d)\n", taskId)
}
