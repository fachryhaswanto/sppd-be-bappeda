package service

import (
	"sppd/model"
	"sppd/repository"
)

type UserService interface {
	FindByUser(username string, password string) (model.User, error)
	Login(username string) (model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) FindByUser(username string, password string) (model.User, error) {
	var user, err = s.userRepository.FindByUser(username, password)

	return user, err
}

func (s *userService) Login(username string) (model.User, error) {
	var user, err = s.userRepository.Login(username)

	return user, err
}
