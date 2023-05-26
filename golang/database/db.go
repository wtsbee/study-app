package database

import (
	"log"
	"mypackage/app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConnect() *gorm.DB {
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "study_app_db"

	dsn := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("gorm接続エラー: ", err)
	}
	log.Println("gorm接続完了")
	return db
}

func Migrate(db *gorm.DB) {
	// db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Article{})
}
