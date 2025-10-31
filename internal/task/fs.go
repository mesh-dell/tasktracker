package task

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
)

// read file
// marshal it to json format

func ReadJsonData() ([]Task, error) {
	filePath := tasksFilePath()
	if !FileExists() {
		fmt.Println("File doesn't exist. Creating file.......")
		file, err := os.Create(filePath)
		os.WriteFile("data.json", []byte("[]"), os.ModeAppend.Perm())

		if err != nil {
			fmt.Println("Error creating file:", err)
			return nil, err
		}

		defer file.Close()
		return []Task{}, nil
	}

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}

	defer file.Close()

	tasks := []Task{}
	byteData, err := io.ReadAll(file)

	if err != nil {
		fmt.Println("Error reading file", err)
		return nil, err
	}

	err = json.Unmarshal(byteData, &tasks)

	if err != nil {
		fmt.Println("Error decoding json", err)
		return nil, err
	}

	return tasks, err
}

func WriteJsonData(tasks []Task) error {
	filePath := tasksFilePath()
	jsonData, err := json.Marshal(tasks)

	if err != nil {
		fmt.Println("Error encoding json", err)
		return err
	}

	err = os.WriteFile(filePath, jsonData, os.ModeAppend.Perm())

	if err != nil {
		fmt.Println("Error writing file", err)
		return err
	}

	return nil
}

func FileExists() bool {
	if _, err := os.Stat("data.json"); err == nil {
		return true
	}
	return false
}

func tasksFilePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		println("Error getting current working directory")
		return ""
	}
	return path.Join(cwd, "data.json")
}
