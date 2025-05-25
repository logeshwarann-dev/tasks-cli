package pkg

import "errors"

func VerifyUserInput(input []string) error {
	if len(input) < 2 {
		return errors.New("invalid command - use add, update, delete and list")
	}

	return nil

}
