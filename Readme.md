### 構成
* go... 言語
* gin... フレームワーク
* golang-migrate... マイグレーション  
* gorm... DBコネクションにのみ使用  
* air... ホットリロード  
* dig... DIに使用

### DB接続(dockerから接続すると思うので不要)
ローカルから
```
psql -h 127.0.0.1 -p 5432 db postgres
```
を実行する

※ psqlコマンドがない場合は以下を実行
```
brew install postgresql@15
brew cleanup postgresql@15
brew linnk postgresql@15
```

### migrationファイルの作成
ローカルから
```
make local-create-migration FILENAME=artist
```

### migrate up
ローカルから
```
make local-migrate-up
```
