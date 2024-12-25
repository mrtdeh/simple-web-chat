package database

import (
	"api-channel/pkg/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*ChatDatabase, error) {
	dsn := "host=localhost user=apiservices password=12345 dbname=test port=5432 sslmode=disable TimeZone=Asia/Tehran"
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database")
	}

	// db.SetupJoinTable(&models.Group{}, "Members", &models.ChatMember{})
	db.AutoMigrate(&models.User{}, &models.Chat{}, &models.ChatMember{},
		&models.Message{}, &models.Attachment{}, &models.Group{},
		models.Thumbnail{}, models.Reply{}, models.ReplyThumbnails{},
		&models.Sticker{})

	return &ChatDatabase{db}, nil

}
