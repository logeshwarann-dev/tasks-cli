package main

import (
	"os"

	"github.com/logeshwarann-dev/taskcli/pkg"
)

func main() {

	if err := pkg.ReadCommandArgs(); err != nil {
		pkg.PrintError(err)
		os.Exit(1)
	}

	// pkg.VerifyUserInput()

}
