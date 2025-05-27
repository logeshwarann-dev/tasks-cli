package pkg

import (
	"flag"
	"fmt"
)

// var CmdReader *bufio.Reader

// func init() {
// 	CmdReader = bufio.NewReader(os.Stdin)
// }

// func ReadInput() error {
// 	userInput, err := CmdReader.ReadString('\n')
// 	if err != nil {
// 		log.Fatal("Error reading user input")
// 		return err
// 	}
// 	userInput = strings.TrimSpace(userInput)
// 	fmt.Println("User Input: ", userInput)
// 	// models.CmdInputSlice = strings.Split(userInput, " ")
// 	return nil
// }

func ReadCommandArgs() ([]string, error) {

	flag.Parse()
	cmdArgs := flag.Args()
	if err := VerifyUserInput(cmdArgs); err != nil {
		return nil, err
	}
	// fmt.Println("Command Args: ", cmdArgs)
	return cmdArgs, nil
}

func PrintError(err error) {
	fmt.Println("Error occurred: ", err.Error())
}
