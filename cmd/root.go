package cmd

import (
	"errors"
	"fmt"
)

func Execute(args []string) error {

	if len(args) < 2 {
		return errors.New("task-cli: No command provided")
	}

	command := args[1]
	switch command {
	case "add":
		return AddCommand(args[2:])
	case "update":
		return UpdateTaskCommand(args[2:])
	case "delete":
		return DeleteCommand(args[2:])
	case "mark-in-progress":
		return MarkInProgressCommand(args[2:])
	case "mark-done":
		return MarkDoneCommand(args[2:])
	case "list":
		return ListCommand(args[2:])
	case "help":
		printHelp()
		return nil
	default:
		return fmt.Errorf("task-cli: unknown command %s", command)
	}
}

func printHelp() {
	fmt.Println(`task-cli — Simple Task Manager

Usage:
  task-cli <command> [arguments]

Commands:
  add "<desc>"            Add a new task
  update <id> "<desc>"    Update an existing task
  delete <id>             Delete a task
  mark-in-progress <id>   Mark a task as in progress
  mark-done <id>          Mark a task as done
  list [status]           List all or filtered tasks (done|todo|in-progress)
  help                    Show this help message

Examples:
  task-cli add "Buy groceries"
  task-cli update 1 "Buy groceries and cook dinner"
  task-cli mark-done 1
  task-cli list done`)
}
