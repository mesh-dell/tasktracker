package cmd

import (
	"fmt"

	"strconv"

	"github.com/mesh-dell/tasktracker/internal/task"
)

func UpdateTaskCommand(args []string) error {
	return RunUpdateTask(args)
}

func MarkInProgressCommand(args []string) error {
	return RunUpdateStatus(args, task.TASK_STATUS_IN_PROGRESS)
}

func MarkDoneCommand(args []string) error {
	return RunUpdateStatus(args, task.TASK_STATUS_DONE)
}

func RunUpdateStatus(args []string, status task.TaskStatus) error {
	if len(args) == 0 {
		return fmt.Errorf("taskId is required")
	}

	taskId, err := strconv.ParseInt(args[0], 10, 64)

	if err != nil {
		return err
	}

	return task.UpdateTaskStatus(taskId, status)
}

func RunUpdateTask(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("please provide a taskId and new description")
	}

	taskId := args[0]
	taskIdInt, err := strconv.ParseInt(taskId, 10, 64)

	if err != nil {
		return err
	}

	newDescription := args[1]
	return task.UpdateTaskDescription(taskIdInt, newDescription)
}
