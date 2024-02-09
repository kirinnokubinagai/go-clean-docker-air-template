package repository

import (
	"go_app/constants"
	"go_app/model"
	"go_app/util"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserList() ([]model.User, error)
}

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) IUserRepository {
	return &UserRepository{database}
}

func (userRepository *UserRepository) GetUserList() ([]model.User, error) {
	var userList []model.User
	sql, err := util.ReadSQLFile(constants.GetUserListSqlPath)
	if err != nil {
		return nil, err
	}

	userRepository.database.Raw(sql).Scan(&userList)
	return userList, nil
}
