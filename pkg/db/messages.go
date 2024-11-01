package database

import (
	"api-channel/pkg/models"
)

func (db *ChatDatabase) GetMessages(chatID uint) ([]models.Message, error) {

	var messages []models.Message
	err := db.gormDB.Where(&models.Message{ChatID: chatID}).
		Preload("Sender").
		Preload("Replies").
		Preload("Replies.ReplyMessage").
		Preload("Replies.Thumbnails").
		Preload("Replies.Thumbnails.Thumbnail").
		Preload("Attachments").
		Preload("Attachments.Thumbnails", "type = ?", "placeholder").
		Find(&messages).Error

	if err != nil {
		return nil, err
	}
	return messages, nil
}
func (db *ChatDatabase) CreateMessage(chatID, userID uint, content string) (uint, error) {
	// Insert message into messages table
	msg := models.Message{
		SenderID: userID,
		ChatID:   chatID,
		Content:  content,
	}
	err := db.gormDB.Create(&msg).Error
	if err != nil {
		return 0, err
	}

	err = db.gormDB.Table("chats").
		Where("id = ?", chatID).
		Update("last_message_id", msg.ID).Error
	if err != nil {
		return 0, err
	}

	return msg.ID, nil
}
