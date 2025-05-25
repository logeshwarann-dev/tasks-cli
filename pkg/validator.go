package pkg

import (
	"errors"

	"github.com/logeshwarann-dev/taskcli/internal/models"
	pkg "github.com/logeshwarann-dev/taskcli/pkg/utils"
)

func VerifyUserInput(input []string) error {
	if len(input) == 0 {
		return errors.New("no command was given.\nuse add, update, delete and list")
	}

	if err := verifyCommand(input[0]); err != nil {
		return err
	}

	switch input[0] {
	case models.AddCommand:
		if err := verifyAddCommand(input); err != nil {
			return err
		}
	case models.ListCommand:
		if err := verfiyListCommand(input); err != nil {
			return err
		}
	case models.MarkDone:
		if err := verifyMarkCommand(input); err != nil {
			return err
		}
	case models.MarkInProgress:
		if err := verifyMarkCommand(input); err != nil {
			return err
		}
	case models.UpdateCommand:
		if err := verifyUpdateCommand(input); err != nil {
			return err
		}
	case models.DeleteCommand:
		if err := verifyDeleteCommand(input); err != nil {
			return err
		}
	default:
		return errors.New("invalid command.\nuse add, update, delete and list")
	}

	return nil

}

func verifyCommand(cmd string) error {
	switch cmd {
	case models.AddCommand:
		break
	case models.ListCommand:
		break
	case models.MarkDone:
		break
	case models.MarkInProgress:
		break
	case models.UpdateCommand:
		break
	case models.DeleteCommand:
		break
	default:
		return errors.New("invalid command was given.\nuse add, update, delete and list")
	}

	return nil
}

func verifyAddCommand(args []string) error {

	if len(args) < 2 {
		return errors.New("invalid args for Add command. \nSyntax: add <\"task name\">")
	}

	if pkg.ValidateNumber(args[1]) {
		return errors.New("invalid args for Add command. \nSyntax: add <\"task name\">")
	}

	if len(args) > 50 {
		return errors.New("task description should be less than 50 words")
	}
	return nil
}

func verfiyListCommand(args []string) error {
	if len(args) > 2 {
		return errors.New("invalid args for List command. \nSyntax: list or list <status>")
	}

	if len(args) == 1 {
		return nil // only list command
	}

	switch args[1] {
	case models.StatusDone:
		break
	case models.StatusInProgress:
		break
	case models.StatusTodo:
		break
	default:
		return errors.New("invalid status for List command. \nuse todo, done, in-progress")
	}

	return nil
}

func verifyDeleteCommand(args []string) error {
	if len(args) > 2 {
		return errors.New("invalid args for Delete command. \nSyntax: delete <id>")
	}

	if len(args) == 1 {
		return errors.New("insufficient args for Delete command. \nSyntax: delete <id>")
	}

	if err := validateTaskId(args[1]); err != nil {
		return err
	}

	return nil
}

func validateTaskId(id string) error {
	if !pkg.ValidateNumber(id) {
		return errors.New("invalid id. \nid should be a positive number")
	}

	if len(id) > 10000 {
		return errors.New("very high id. \nid should be valid")
	}

	return nil
}

func verifyMarkCommand(args []string) error {
	if len(args) > 2 {
		return errors.New("invalid args for Mark command. \nSyntax: mark-done or mark-in-progress <id>")
	}

	if len(args) == 1 {
		return errors.New("insufficient args for Mark command. \nSyntax: mark-done or mark-in-progress <id>")
	}

	if err := validateTaskId(args[1]); err != nil {
		return err
	}

	return nil
}

func verifyUpdateCommand(args []string) error {
	if len(args) < 2 {
		return errors.New("invalid args for Update command. \nSyntax: update <id> <\"task name\">")
	}

	if !pkg.ValidateNumber(args[1]) {
		return errors.New("invalid id for Update command. \nSyntax: update <id> <\"task name\">")
	}

	if len(args) > 50 {
		return errors.New("task description should be less than 50 words")
	}
	return nil
}
