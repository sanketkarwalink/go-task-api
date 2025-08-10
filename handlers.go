package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func handleTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tasks, err := loadTasks()
		if err != nil {
			http.Error(w, "Could not load tasks", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
		return
	}

	if r.Method == http.MethodPost {
		tasks, err := loadTasks()
		if err != nil {
			http.Error(w, "Could not load tasks", http.StatusInternalServerError)
			return
		}

		var newTask Task
		err = json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		maxID := 0
		for _, t := range tasks {
			if t.ID > maxID {
				maxID = t.ID
			}
		}
		newTask.ID = maxID + 1

		tasks = append(tasks, newTask)
		err = saveTasks(tasks)
		if err != nil {
			http.Error(w, "Could not save task", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newTask)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func handleGetTaskByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/tasks/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	tasks, err := loadTasks()
	if err != nil {
		http.Error(w, "Could not load tasks", http.StatusInternalServerError)
		return
	}

	for _, t := range tasks {
		if t.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(t)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func handleUpdateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/tasks/update/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	tasks, err := loadTasks()
	if err != nil {
		http.Error(w, "Could not load tasks", http.StatusInternalServerError)
		return
	}

	var updateTask Task
	err = json.NewDecoder(r.Body).Decode(&updateTask)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updated := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Name = updateTask.Name
			tasks[i].Done = updateTask.Done
			updated = true
			break
		}
	}

	if !updated {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	err = saveTasks(tasks)
	if err != nil {
		http.Error(w, "Could not save updated task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Task updated successfully"})
}

func handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/tasks/delete/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	tasks, err := loadTasks()
	if err != nil {
		http.Error(w, "Could not load tasks", http.StatusInternalServerError)
		return
	}

	found := false
	newTasks := []Task{}
	for _, t := range tasks {
		if t.ID != id {
			newTasks = append(newTasks, t)
		} else {
			found = true
		}
	}

	if !found {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	err = saveTasks(newTasks)
	if err != nil {
		http.Error(w, "Could not save updated task list", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
}
