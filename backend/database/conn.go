package database

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IDbConnection interface {
	// reader用db接続
	ConnectReaderDb()

	// writer用db接続
	ConnectWriterDb()

	// reader用dbコネクションを返す
	GetReader() *gorm.DB

	// writer用dbコネクションを返す
	GetWriter() *gorm.DB
}

type DbConnection struct {
	// reader用のコネクション
	Reader *gorm.DB

	// writer用のコネクション
	Writer *gorm.DB
}

// コンストラクター
// @return IDbConnection インターフェース
func NewDbConnection() IDbConnection {
	dbConn := &DbConnection{}
	dbConn.ConnectReaderDb()
	dbConn.ConnectWriterDb()
	return dbConn
}

// reader用db接続
func (dbConn *DbConnection) ConnectReaderDb() {
	host := os.Getenv("POSTGRES_READER_HOST")
	user := os.Getenv("POSTGRES_READER_USER")
	password := os.Getenv("POSTGRES_READER_PASSWORD")
	dbName := os.Getenv("POSTGRES_READER_DB")
	sslMode := os.Getenv("POSTGRES_READER_SSL_MODE")
	TimeZone := os.Getenv("POSTGRES_READER_TIME_ZONE")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=%v TimeZone=%v", host, user, password, dbName, sslMode, TimeZone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalln(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logrus.Fatalln(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	dbConn.Reader = db
}

// writer用db接続
func (dbConn *DbConnection) ConnectWriterDb() {
	host := os.Getenv("POSTGRES_WRITER_HOST")
	user := os.Getenv("POSTGRES_WRITER_USER")
	password := os.Getenv("POSTGRES_WRITER_PASSWORD")
	dbName := os.Getenv("POSTGRES_WRITER_DB")
	sslMode := os.Getenv("POSTGRES_WRITER_SSL_MODE")
	TimeZone := os.Getenv("POSTGRES_WRITER_TIME_ZONE")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=%v TimeZone=%v", host, user, password, dbName, sslMode, TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalln(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logrus.Fatalln(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	dbConn.Writer = db
}

// reader用のコネクションを返す
// @return *gorm.DB reader用のコネクション
func (dbConn *DbConnection) GetReader() *gorm.DB {
	return dbConn.Reader
}

// writer用のコネクションを返す
// @return *gorm.DB writer用のコネクション
func (dbConn *DbConnection) GetWriter() *gorm.DB {
	return dbConn.Writer
}
