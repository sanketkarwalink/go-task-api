# 📝 Go Task API

![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)

A simple and clean **RESTful Task Manager API** built with **Go**.  
Supports creating, reading, updating, and deleting tasks stored in a local JSON file.  
Perfect for learning **Go backend development** and demonstrating API skills.

---

## ✨ Features

- ➕ **Create** a new task  
- 📋 **Get all** tasks  
- 🔍 **Get** a task by ID  
- ✏ **Update** a task by ID  
- ❌ **Delete** a task by ID  
- 💾 Stores tasks in a `tasks.json` file (no external DB required)  

---

## 📂 Project Structure

```

go-task-api/
│
├── main.go        # Entry point, starts server, sets up routes
├── handlers.go    # HTTP request handlers
├── models.go      # Data models (Task struct)
├── storage.go     # File I/O for tasks.json
├── tasks.json     # Local storage for tasks
├── go.mod
├── go.sum
└── README.md

```

---

## 🔌 API Endpoints

### 1. Get all tasks
```

GET /tasks

```

### 2. Get a task by ID
```

GET /tasks/{id}

```

### 3. Create a task
```

POST /tasks
Body: { "name": "Task Name", "done": false }

```

### 4. Update a task
```

PUT /tasks/update/{id}
Body: { "name": "Updated Task", "done": true }

```

### 5. Delete a task
```

DELETE /tasks/delete/{id}

````

---

## 🚀 Getting Started

### 1️⃣ Clone the repo
```bash
git clone https://github.com/sanketkarwalink/go-task-api.git
cd go-task-api
````

### 2️⃣ Install dependencies

```bash
go mod tidy
```

### 3️⃣ Run the server

```bash
go run main.go handlers.go models.go storage.go
```

Server will start at:

```
http://localhost:8080
```

---

## 📌 Why This Project?

This project demonstrates:

* ✅ Backend API development with Go
* ✅ REST API design principles
* ✅ File-based storage handling
* ✅ Modular code organization
* ✅ JSON data handling in Go

---

## 🤝 Contributing

Pull requests are welcome!
For major changes, please open an issue first to discuss what you would like to change.

---

## 📜 License

This project is licensed under the [MIT License](LICENSE) — feel free to use it.

---

**👨‍💻 Developed by [Sanket Karwa](https://github.com/sanketkarwalink)**
