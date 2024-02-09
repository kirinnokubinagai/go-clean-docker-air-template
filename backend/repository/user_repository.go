package repository

import (
	constants "go_app/constant"
	"go_app/database"
	"go_app/model"
	"go_app/util"
)

type IUserRepository interface {
	// ユーザー一覧取得処理
	GetUserList() ([]model.User, error)
}

type UserRepository struct {
	// DBクライアント
	database database.IDbConnection
}

// コンストラクター
// @param database database.IDbConnection DBクライアント
// @return IUserRepository インターフェース
func NewUserRepository(database database.IDbConnection) IUserRepository {
	return &UserRepository{database}
}

// ユーザー一覧取得処理
// @return []model.User ユーザー情報一覧
// @return error エラー
func (userRepository *UserRepository) GetUserList() ([]model.User, error) {
	var userList []model.User
	sql, err := util.ReadSQLFile(constants.GetUserListSqlPath)
	if err != nil {
		return nil, err
	}

	userRepository.database.GetWriter().Raw(sql).Scan(&userList)
	return userList, nil
}
