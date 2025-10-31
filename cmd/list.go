package cmd

import "github.com/mesh-dell/tasktracker/internal/task"

func ListCommand(args []string) error {

	if len(args) > 0 {
		status := task.TaskStatus(args[0])
		return task.ListTasks(status)
	}

	return task.ListTasks("all")
}
