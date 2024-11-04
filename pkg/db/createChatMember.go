package database

import (
	"api-channel/pkg/models"
	"time"
)

func (db *ChatDatabase) CreateChatMember(chatID uint32, userID uint32, role string) error {
	now := time.Now()

	member := &models.ChatMember{
		ChatID:   chatID,
		UserID:   userID,
		Role:     role,
		JoinedAt: now,
	}
	err := db.gormDB.Create(member).Error
	if err != nil {
		return err
	}

	return nil
}
