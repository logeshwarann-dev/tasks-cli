package app

import (
	"fmt"
	"log"
	"os"

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
		fmt.Println("Inside Add Command case")
		taskDesc := utils.ConvSliceToStr(CommandArgs[1:])
		if err := services.AddTask(taskDesc); err != nil {
			pkg.PrintError(err)
			os.Exit(1)
		}
	default:
		log.Fatal("Invalid command")
	}

}
