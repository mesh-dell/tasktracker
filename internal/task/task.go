package task

import (
	"fmt"
	"time"
)

type TaskStatus string

const (
	TASK_STATUS_TODO        TaskStatus = "todo"
	TASK_STATUS_IN_PROGRESS TaskStatus = "in-progress"
	TASK_STATUS_DONE        TaskStatus = "done"
)

type Task struct {
	Id          int
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTask(id int, description string) *Task {
	var task = Task{
		Id:          id,
		Description: description,
		Status:      TASK_STATUS_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return &task
}

func AddTask(description string) error {
	existingTasks, err := ReadJsonData()

	if err != nil {
		return err
	}

	var newTaskId int

	if len(existingTasks) == 0 {
		newTaskId = 1
	} else {
		lastTask := existingTasks[len(existingTasks)-1]
		newTaskId = lastTask.Id + 1
	}

	newTask := NewTask(newTaskId, description)
	existingTasks = append(existingTasks, *newTask)
	fmt.Println("Added task successfully")
	return WriteJsonData(existingTasks)
}

func UpdateTaskStatus(id int64, status TaskStatus) error {
	allTasks, err := ReadJsonData()

	if err != nil {
		return err
	}

	var taskExists bool
	var updatedTasks []Task

	for _, task := range allTasks {
		if task.Id == int(id) {
			taskExists = true
			task.Status = status
			task.UpdatedAt = time.Now()
		}
		updatedTasks = append(updatedTasks, task)
	}

	if !taskExists {
		return fmt.Errorf("task not found(ID: %d)", id)
	}

	fmt.Printf("\nTask updated successfully: %d\n\n", id)
	return WriteJsonData(updatedTasks)
}

func UpdateTaskDescription(id int64, description string) error {
	allTasks, err := ReadJsonData()

	if err != nil {
		return err
	}

	var taskExists bool
	var updatedTasks []Task

	for _, task := range allTasks {
		if task.Id == int(id) {
			taskExists = true
			task.Description = description
			task.UpdatedAt = time.Now()
		}
		updatedTasks = append(updatedTasks, task)
	}

	if !taskExists {
		return fmt.Errorf("task not found(ID: %d)", id)
	}
	fmt.Printf("\nTask updated successfully: %d\n\n", id)
	return WriteJsonData(updatedTasks)
}

func DeleteTask(id int64) error {
	allTasks, err := ReadJsonData()
	if err != nil {
		return err
	}

	var updatedTasks []Task

	for _, task := range allTasks {
		if task.Id != int(id) {
			updatedTasks = append(updatedTasks, task)
		}
	}

	if len(updatedTasks) == len(allTasks) {
		return fmt.Errorf("task not found(ID: %d)", id)
	}

	fmt.Printf("\nTask deleted successfully: %d\n\n", id)
	return WriteJsonData(updatedTasks)
}

func ListTasks(status TaskStatus) error {
	tasks, err := ReadJsonData()

	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	var filteredTasks []Task

	switch status {
	case "all":
		filteredTasks = tasks
	case TASK_STATUS_DONE:
		for _, task := range tasks {
			if task.Status == TASK_STATUS_DONE {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case TASK_STATUS_IN_PROGRESS:
		for _, task := range tasks {
			if task.Status == TASK_STATUS_IN_PROGRESS {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case TASK_STATUS_TODO:
		for _, task := range tasks {
			if task.Status == TASK_STATUS_TODO {
				filteredTasks = append(filteredTasks, task)
			}
		}
	}

	fmt.Println()

	if len(filteredTasks) == 0 {
		fmt.Printf("There are no %s tasks!\n", status)
		return nil
	}

	fmt.Printf("Here are your %s tasks:\n\n", status)
	for _, task := range filteredTasks {
		relativeUpdatedTime := task.UpdatedAt.Format("2006-01-02 15:04:05")
		relativeCreatedTime := task.CreatedAt.Format("2006-01-02 15:04:05")
		fmt.Println("ID:", task.Id)
		fmt.Println("Description:", task.Description)
		fmt.Println("Status:", task.Status)
		fmt.Println("CreatedAt:", relativeCreatedTime)
		fmt.Println("UpdatedAt:", relativeUpdatedTime)
		fmt.Println()
	}
	return nil
}
