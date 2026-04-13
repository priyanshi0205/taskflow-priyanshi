package user

import (
	"backedn_go/model"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) ListExceptUser(userID string) ([]model.User, error) {
	var users []model.User
	if err := r.db.Model(&model.User{}).
		Select("id", "name").
		Order("name ASC").
		Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
