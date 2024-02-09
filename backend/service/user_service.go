package service

import (
	"go_app/model"
	"go_app/repository"
)

type IUserService interface {
	// ユーザー一覧取得処理
	GetUserList() ([]model.User, error)
}

type UserService struct {
	// UserRepositoryのインターフェース
	userRepository repository.IUserRepository
}

// コンストラクター
// @param userRepository repository.IUserRepository ユーザーレポジトリー
// @return IUserService インターフェース
func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{userRepository}
}

/*
ユーザー一覧取得処理
@return []model.User ユーザー情報一覧
@return error エラー
*/
func (userService UserService) GetUserList() ([]model.User, error) {
	return userService.userRepository.GetUserList()
}
