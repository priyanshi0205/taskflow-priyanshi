package model

import "time"

type Task struct {
	ID          string     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title       string     `gorm:"type:text;not null" json:"title"`
	Description string     `gorm:"type:text" json:"description"`
	Status      string     `gorm:"type:text;default:todo" json:"status"`
	Priority    string     `gorm:"type:text;default:medium" json:"priority"`
	ProjectID   string     `gorm:"type:uuid;column:project_id" json:"project_id"`
	CreatedByID string     `gorm:"type:uuid;column:created_by_id" json:"created_by_id"`
	AssigneeID  *string    `gorm:"type:uuid;column:assignee_id" json:"assignee_id,omitempty"`
	DueDate     *time.Time `gorm:"type:date;column:due_date" json:"due_date,omitempty"`
	CreatedAt   time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Task) TableName() string {
	return "tasks"
}
