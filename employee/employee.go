package employee

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Employee struct {
	ID         string         `json:"id"`
	Email      string         `json:"email"`
	Phone      string         `json:"phone"`
	AssignedTo pq.StringArray `json:"assigned_to" gorm:"type:text[]"`
}

func New(email string, phone string) (Employee, error) {
	return Employee{
		ID:         uuid.New().String(),
		Email:      email,
		Phone:      phone,
		AssignedTo: []string{},
	}, nil
}
