package database

import (
	"api-channel/pkg/models"
	"fmt"
	"time"
)

func (db *ChatDatabase) GetMessages(chatID, msgID, nextCount, prevCount uint32) ([]models.Message, error) {
	time.Sleep(time.Millisecond * 700)
	// var messages []models.Message
	fmt.Println("get messages:", msgID, prevCount, int(msgID)-int(prevCount))
	fromMsgID := 1
	toMsgID := 50
	if nextCount > 0 {
		fromMsgID = int(msgID)
		toMsgID = int(msgID + nextCount)
	}
	if prevCount > 0 {

		fromMsgID = int(msgID)
		toMsgID = int(msgID - prevCount)
		if toMsgID < 1 {
			toMsgID = 1
		}
	}
	// newFrom := int(msgID) - int(prevCount)
	// if newFrom > 0 {
	// 	fromMsgID = newFrom
	// }
	// toMsgID := msgID + nextCount
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
	// }

	// Lazy-Loading prev messages
	// prevMessages, err := fetchMessages("chat_id = ? AND id < ?", "id DESC", int(prevCount))
	// if err != nil {
	// 	return nil, err
	// }
	// Lazy-Loading next messages
	// nextMessages, err := fetchMessages("chat_id = ? AND id+? < ?", "id ASC", int(nextCount))
	// if err != nil {
	// 	return nil, err
	// }
	// Merge all messages
	// messages = nextMessages //append(prevMessages, nextMessages...)
	// return messages, nil
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
