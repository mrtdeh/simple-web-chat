package database

import (
	"api-channel/pkg/models"
)

func (db *ChatDatabase) CreateChat(isGroup bool) (uint32, error) {

	// Insert message into messages table
	chat := models.Chat{
		IsGroup: isGroup,
	}
	err := db.gormDB.Create(&chat).Error
	if err != nil {
		return 0, err
	}

	return uint32(chat.ID), nil
}
