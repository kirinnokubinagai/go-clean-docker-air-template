package repository

import (
	"fmt"
	"go_app/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserList() []model.User
}

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) IUserRepository {
	return &UserRepository{database}
}

func (userRepository *UserRepository) GetUserList() []model.User {
	var userList []model.User
	userRepository.database.Raw(`
		SELECT 
			*
	  FROM
			"user"
 `).Scan(&userList)
	fmt.Println(userList)
	return userList
}
