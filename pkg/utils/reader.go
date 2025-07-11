package utils

import (
	"encoding/json"
	"os"

	"github.com/logeshwarann-dev/taskcli/internal/models"
)

func IsFileEmpty(filePath string) (bool, error) {
	tasksData, err := os.ReadFile(filePath)
	if err != nil {
		return false, err
	}
	if len(tasksData) == 0 {
		return true, nil
	}
	return false, nil
}

func ReadJSONFile(filePath string, tasks *models.Tasks) error {
	tasksData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(tasksData), &tasks); err != nil {
		return err
	}

	return nil
}

func CheckIfFileExists(filePath string) bool {

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return false
	}
	defer file.Close()
	return true

}
