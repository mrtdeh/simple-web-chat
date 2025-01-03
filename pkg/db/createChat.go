package database

import (
	"api-channel/pkg/models"
)

func (db *ChatDatabase) CreateChat(chatType string) (uint32, error) {

	// Insert message into messages table
	chat := models.Chat{
		Type: chatType,
	}
	err := db.gormDB.Create(&chat).Error
	if err != nil {
		return 0, err
	}

	return uint32(chat.ID), nil
}
