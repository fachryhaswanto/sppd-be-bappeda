package service

import (
	"sppd/helper"
	"sppd/model"
	"sppd/repository"
	"sppd/request"
)

type UserService interface {
	FindAll() ([]model.User, error)
	FindByUser(username string, password string) (model.User, error)
	Login(username string) (model.User, error)
	Update(id int, userRequest request.UpdateUserRequest) (model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) FindAll() ([]model.User, error) {
	var users, err = s.userRepository.FindAll()

	return users, err
}

func (s *userService) FindByUser(username string, password string) (model.User, error) {
	var user, err = s.userRepository.FindByUser(username, password)

	return user, err
}

func (s *userService) Login(username string) (model.User, error) {
	var user, err = s.userRepository.Login(username)

	return user, err
}

func (s *userService) Update(id int, userRequest request.UpdateUserRequest) (model.User, error) {
	var user, err = s.userRepository.FindById(id)

	if userRequest.Password != "" {
		var password = userRequest.Password
		var hashedPassword, _ = helper.HashPassword(password)
		userRequest.Password = hashedPassword
	}

	user.Username = userRequest.Username
	user.Password = userRequest.Password
	user.Role = userRequest.Role

	updatedUser, err := s.userRepository.Update(user)

	return updatedUser, err
}
