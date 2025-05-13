package internal

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	todo       = "todo"
	inProgress = "in-progress"
	done       = "done"
)

type Task struct {
	Id          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewTask(description string) (*Task, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("Error al instanciar el UUID: %w", err)
	}

	task := &Task{
		Id:          id,
		Description: description,
		Status:      todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return task, nil
}
