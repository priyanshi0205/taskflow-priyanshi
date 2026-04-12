package model

import "time"

type User struct {
	ID        string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string    `gorm:"type:text;not null" json:"name"`
	Email     string    `gorm:"type:text;uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"type:text;not null" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (User) TableName() string {
	return "users"
}
