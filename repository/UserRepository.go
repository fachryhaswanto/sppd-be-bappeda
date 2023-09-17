package repository

import (
	"sppd/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindById(id int) (model.User, error)
	FindByUser(username string, password string) (model.User, error)
	Login(username string) (model.User, error)
	Update(user model.User) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAll() ([]model.User, error) {
	var users []model.User

	var err = r.db.Model(&users).Find(&users).Error

	return users, err
}

func (r *userRepository) FindById(id int) (model.User, error) {
	var user model.User
	var err = r.db.Model(&user).Take(&user, id).Error

	return user, err
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

func (r *userRepository) Update(user model.User) (model.User, error) {
	var err = r.db.Model(&user).Updates(model.User{
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
	}).Error

	return user, err
}
