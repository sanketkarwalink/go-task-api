package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/tasks", handleTasks)
	http.HandleFunc("/tasks/", handleGetTaskByID)
	http.HandleFunc("/tasks/update/", handleUpdateTask)
	http.HandleFunc("/tasks/delete/", handleDeleteTask)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
