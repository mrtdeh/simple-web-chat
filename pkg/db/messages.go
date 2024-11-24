package database

import (
	"api-channel/pkg/models"
)

func (db *ChatDatabase) GetMessages(chatID, msgID, nextCount, prevCount uint32) ([]models.Message, error) {
	var messages []models.Message

	fetchMessages := func(condition string, order string, limit int) ([]models.Message, error) {
		var result []models.Message
		err := db.gormDB.Where(condition, chatID, msgID).
			Order(order).
			Limit(limit).
			Preload("Sender").
			Preload("Replies").
			Preload("Replies.ReplyMessage").
			Preload("Replies.Thumbnails").
			Preload("Replies.Thumbnails.Thumbnail").
			Preload("Attachments").
			Preload("Attachments.Thumbnails", "type = ?", "placeholder").
			Find(&result).Error
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	// Lazy-Loading prev messages
	prevMessages, err := fetchMessages("chat_id = ? AND id < ?", "id DESC", int(prevCount))
	if err != nil {
		return nil, err
	}
	// Lazy-Loading next messages
	nextMessages, err := fetchMessages("chat_id = ? AND id > ?", "id ASC", int(nextCount))
	if err != nil {
		return nil, err
	}
	// Merge all messages
	messages = append(prevMessages, nextMessages...)
	return messages, nil
}

func (db *ChatDatabase) CreateMessage(chatID, userID uint32, content string) (uint32, error) {
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
