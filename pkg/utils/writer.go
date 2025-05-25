package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/logeshwarann-dev/taskcli/internal/models"
)

func WriteToFile(filePath string, tasksList models.Tasks) error {

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error while opening file: %v", err)
	}
	defer file.Close()

	tasksByteArr, err := json.Marshal(tasksList)
	if err != nil {
		return fmt.Errorf("error while marshalling tasks: %v", err)
	}

	_, err = file.Write(tasksByteArr)
	if err != nil {
		return fmt.Errorf("error while writing data to file: %v", err)
	}

	return nil

}
