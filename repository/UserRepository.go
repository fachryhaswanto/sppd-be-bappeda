package repository

import (
	"sppd/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUser(username string, password string) (model.User, error)
	Login(username string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByUser(username string, password string) (model.User, error) {
	var user model.User

	var err = r.db.Where(model.User{Username: username, Password: password}).Take(&user).Error

	return user, err
}

func (r *userRepository) Login(username string) (model.User, error) {
	var user model.User

	var err = r.db.Where("username = ?", username).Take(&user).Error

	return user, err
}
