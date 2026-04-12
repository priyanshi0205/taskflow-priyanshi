package model

import "time"

type Project struct {
	ID          string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string    `gorm:"type:text;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	OwnerID     string    `gorm:"type:uuid;column:owner_id" json:"owner_id"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	Tasks       []Task    `gorm:"foreignKey:ProjectID" json:"tasks,omitempty"`
}

func (Project) TableName() string {
	return "projects"
}
