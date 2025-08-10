package main

import (
	"encoding/json"
	"os"
)

func loadTasks() ([]Task, error) {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil // No file yet, return empty list
		}
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}
