![go](https://img.shields.io/github/go-mod/go-version/kerunaru/task-tracker) ![license](https://img.shields.io/github/license/kerunaru/task-tracker)

# TASK TRACKER

## Usage

As simple as:

```bash
# Insert a new task
go run cmd/main.go add "Task description"

# Update an existing task
go run cmd/main.go update 638925e2-2fc6-11f0-b5c2-2800af748d54 "Updated task description"

# Delete a task
go run cmd/main.go delete 638925e2-2fc6-11f0-b5c2-2800af748d54

# List all tasks
go run cmd/main.go list

# List all tasks by status
go run cmd/main.go list done
go run cmd/main.go list todo
go run cmd/main.go list in-progress

# Change task status
go run cmd/main.go mark done 638925e2-2fc6-11f0-b5c2-2800af748d54
go run cmd/main.go mark in-progress 638925e2-2fc6-11f0-b5c2-2800af748d54
go run cmd/main.go mark todo 638925e2-2fc6-11f0-b5c2-2800af748d54
```

This is a Go solution of [roadmap.sh Task Tracker exercise](https://roadmap.sh/projects/task-tracker)
