package service

import (
	"go_app/model"
	"go_app/repository"
)

type IUserService interface {
	GetUserList() []model.User
}

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{userRepository}
}

func (userService UserService) GetUserList() []model.User {
	return userService.userRepository.GetUserList()
}
