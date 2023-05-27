package database

import (
	"fmt"
	"log"
	"os"
	"study-app-api/model"

	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Println("database NewDB 環境変数読込みエラー: ", err)
	}

	USER := os.Getenv("MYSQL_ROOT_USER")
	PASS := os.Getenv("MYSQL_ROOT_PASSWORD")
	DBNAME := os.Getenv("MYSQL_DATABASE")
	PROTOCOL := fmt.Sprintf("tcp(db:%s)", os.Getenv("PORT"))

	dsn := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database NewDB DB接続エラー: ", err)
	}
	log.Println("database NewDB DB接続完了")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Println("database CloseDB DBクローズエラー: ", err)
	}
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
