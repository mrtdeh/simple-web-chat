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

func (db *ChatDatabase) SetChatLastReadedMessageID(chatID, userID, lastReadedMessageID uint32) error {
	var member models.ChatMember
	err := db.gormDB.Where("chat_id = ? and user_id = ?", chatID, userID).First(&member).Error
	if err != nil {
		return err
	}
	member.LastReadedMessageID = lastReadedMessageID
	err = db.gormDB.Save(&member).Error
	if err != nil {
		return err
	}

	return nil
}

func (db *ChatDatabase) GetChatLastReadedMessageID(chatID, userID uint32) (uint32, error) {
	var member models.ChatMember
	err := db.gormDB.Where("chat_id = ? and user_id = ?", chatID, userID).First(&member).Error
	if err != nil {
		return 0, err
	}
	return member.LastReadedMessageID, nil
}
