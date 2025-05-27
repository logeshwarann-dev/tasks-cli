package app

import (
	"log"
	"os"
	"strconv"

	"github.com/logeshwarann-dev/taskcli/internal/models"
	services "github.com/logeshwarann-dev/taskcli/internal/services/crud"
	"github.com/logeshwarann-dev/taskcli/pkg"
	"github.com/logeshwarann-dev/taskcli/pkg/utils"
)

func Start() {

	var CommandArgs []string
	var err error
	if CommandArgs, err = pkg.ReadCommandArgs(); err != nil {
		pkg.PrintError(err)
		os.Exit(1)
	}

	switch CommandArgs[0] {
	case models.AddCommand:
		taskDesc := utils.ConvSliceToStr(CommandArgs[1:])
		taskId, err := services.AddTask(taskDesc, models.FilePath)
		if err != nil {
			pkg.PrintError(err)
			os.Exit(1)
		}
		services.PrintAddedTask(taskId)
	case models.UpdateCommand:
		taskId, err := strconv.Atoi(CommandArgs[1])
		if err != nil {
			pkg.PrintError(err)
			os.Exit(1)
		}
		taskDesc := utils.ConvSliceToStr(CommandArgs[2:])
		if err := services.UpdateTask(taskId, taskDesc, models.FilePath); err != nil {
			pkg.PrintError(err)
			os.Exit(1)
		}
		services.PrintUpdatedTask(taskId)
	case models.DeleteCommand:
		taskId, err := strconv.Atoi(CommandArgs[1])
		if err != nil {
			pkg.PrintError(err)
			os.Exit(1)
		}
		if err := services.DeleteTask(taskId, models.FilePath); err != nil {
			pkg.PrintError(err)
			os.Exit(1)
		}
		services.PrintDeletedTask(taskId)
	case models.MarkDone:
		taskId, err := strconv.Atoi(CommandArgs[1])
		if err != nil {
			pkg.PrintError(err)
			os.Exit(1)
		}
		if err := services.UpdateTask(taskId, models.StatusDone, models.FilePath); err != nil {
			pkg.PrintError(err)
			os.Exit(1)
		}
		services.PrintTaskStatus(taskId)
	case models.MarkInProgress:
		taskId, err := strconv.Atoi(CommandArgs[1])
		if err != nil {
			pkg.PrintError(err)
			os.Exit(1)
		}
		if err := services.UpdateTask(taskId, models.StatusInProgress, models.FilePath); err != nil {
			pkg.PrintError(err)
			os.Exit(1)
		}
		services.PrintTaskStatus(taskId)
	case models.ListCommand:
		status := services.GetStatusFlag(CommandArgs)
		if err := services.ListTask(status, models.FilePath); err != nil {
			pkg.PrintError(err)
			os.Exit(1)
		}
	default:
		log.Fatal("Invalid command")
	}

}
