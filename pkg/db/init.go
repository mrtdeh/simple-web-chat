package database

import (
	"api-channel/pkg/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() (*gorm.DB, error) {
	dsn := "apiservices:F@rz@n@2022@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database")
	}

	db.SetupJoinTable(&models.Group{}, "Members", &models.ChatMember{})
	db.AutoMigrate(&models.User{}, &models.Chat{},
		&models.Message{}, &models.Attachment{}, &models.Group{},
		models.Thumbnail{}, models.Reply{}, models.ReplyThumbnails{},
		&models.PrivateChat{}, &models.LastMessageRead{}, &models.Sticker{})

	return db, nil

}

// s dasdasdasdasd
func GetInstance() *gorm.DB {
	if db == nil {
		var err error
		db, err = Init()
		if err != nil {
			log.Fatal(err)
		}
	}
	return db
}
