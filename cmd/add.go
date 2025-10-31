package cmd

import (
	"fmt"

	"github.com/mesh-dell/tasktracker/internal/task"
)

func AddCommand(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("task-cli Not enough arguments provided")
	}
	if len(args) > 1 {
		return fmt.Errorf("task-cli Add takes one argument")
	}

	taskDescription := args[0]
	return task.AddTask(taskDescription)
}
