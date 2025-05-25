package app

import (
	"os"

	"github.com/logeshwarann-dev/taskcli/pkg"
)

func Start() {
	if err := pkg.ReadCommandArgs(); err != nil {
		pkg.PrintError(err)
		os.Exit(1)
	}
}
