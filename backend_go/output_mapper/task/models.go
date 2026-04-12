package task

import "time"

type TaskResponse struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Priority    string     `json:"priority"`
	ProjectID   string     `json:"project_id"`
	CreatedByID string     `json:"created_by_id"`
	AssigneeID  *string    `json:"assignee_id,omitempty"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
