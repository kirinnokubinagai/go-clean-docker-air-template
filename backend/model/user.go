package model

import "time"

// ユーザー用のモデル
type User struct {
	// ユーザーID
	UserId uint64 `json:"user_id"`

	// メールアドレス
	Email string `json:"email"`

	// パスワード
	Password string `json:"password"`

	// 作成日
	CreatedAt time.Time `json:"created_at"`

	// 削除日 (ソフトデリート)
	DeletedAt time.Time `json:"deleted_at"`
}
