CREATE TABLE "user"
(
  user_id BIGSERIAL PRIMARY KEY,
  name VARCHAR(64) NOT NULL UNIQUE,
  email VARCHAR(100) NOT NULL UNIQUE,
  password text NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE "user" IS 'ユーザーテーブル';

COMMENT ON COLUMN "user".user_id IS 'ユーザーID';
COMMENT ON COLUMN "user".name IS 'ユーザー名';
COMMENT ON COLUMN "user".email IS 'メールアドレス';
COMMENT ON COLUMN "user".created_at IS '登録日時';
COMMENT ON COLUMN "user".deleted_at IS '削除日時';