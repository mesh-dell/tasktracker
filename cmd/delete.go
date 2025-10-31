package cmd

import (
	"fmt"
	"strconv"

	"github.com/mesh-dell/tasktracker/internal/task"
)

func DeleteCommand(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("taskId is required")
	}

	taskId := args[0]
	taskIdInt, err := strconv.ParseInt(taskId, 10, 64)

	if err != nil {
		return err
	}

	return task.DeleteTask(taskIdInt)

}
