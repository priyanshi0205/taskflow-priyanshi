package task

import "time"

type CreateTaskInput struct {
	Title       string
	Description string
	Status      string
	Priority    string
	CreatedByID string
	AssigneeID  *string
	DueDate     *time.Time
}

type UpdateTaskInput struct {
	Title       *string
	Description *string
	Status      *string
	Priority    *string
	AssigneeID  *string
	DueDate     *time.Time
}
