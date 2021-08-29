package task

import (
	"github.com/google/uuid"
)

type Task struct {
	ID        string `json:"id"`
	Name      string `json:"name" gorm:"unique"`
	Completed bool   `json:"completed"`
}

func New(name string) (Task, error) {
	return Task{
		ID:        uuid.New().String(),
		Name:      name,
		Completed: false,
	}, nil
}
