package database

import (
	"api-channel/pkg/models"
	"fmt"
)

func (db *ChatDatabase) GetMessages(chatID, msgID, nextCount, prevCount uint32) ([]models.Message, error) {
	var fromMsgID int
	var toMsgID int
	if prevCount > 0 && nextCount > 0 {
		// اگر هر دو مقدار وجود داشته باشند
		fromMsgID = int(msgID) - int(prevCount)
		if fromMsgID < 1 {
			fromMsgID = 1 // جلوگیری از مقدار منفی
		}
		toMsgID = int(msgID) + int(nextCount)
	} else if nextCount > 0 {
		// فقط پیام‌های جدیدتر (next)
		fromMsgID = int(msgID)
		toMsgID = int(msgID) + int(nextCount)
	} else if prevCount > 0 {
		// فقط پیام‌های قدیمی‌تر (previous)
		toMsgID = int(msgID)
		fromMsgID = int(msgID) - int(prevCount)
		if fromMsgID < 1 {
			fromMsgID = 1 // جلوگیری از مقدار منفی
		}
	} else {
		// هیچ داده‌ای برای بازه مشخص نشده است
		return nil, fmt.Errorf("no valid range specified")
	}

	var result []models.Message
	err := db.gormDB.Debug().
		Where("chat_id = ? AND id >= ? AND id <= ?", chatID, fromMsgID, toMsgID).
		Order("id ASC").
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
