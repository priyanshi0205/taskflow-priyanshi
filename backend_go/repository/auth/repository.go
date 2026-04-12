package auth

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

func (r *repository) CreateUser(name, email, hashedPassword string) (*model.User, error) {
	user := &model.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
