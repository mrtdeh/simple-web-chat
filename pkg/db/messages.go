package database

import (
	"api-channel/pkg/models"
	"fmt"
	"sort"

	"gorm.io/gorm"
)

func (db *ChatDatabase) GetMessages(chatID, fromMsgID uint32, direction string, count int32) ([]models.Message, error) {
	var result []models.Message

	// Base queries for next and previous
	nextQuery := db.gormDB.Table("messages").Select("*").
		Where("chat_id = ? AND id > ?", chatID, fromMsgID).
		Order("id ASC").
		Limit(int(count))

	previousQuery := db.gormDB.Table("messages").Select("*").
		Where("chat_id = ? AND id < ?", chatID, fromMsgID).
		Order("id DESC").
		Limit(int(count))

	var finalQuery *gorm.DB

	switch direction {
	case "RecordDirection.next":
		// Only next query
		finalQuery = nextQuery
	case "RecordDirection.previous":
		// Only previous query
		finalQuery = previousQuery
	case "RecordDirection.none":
		// Combine both with UNION
		// finalQuery = db.gormDB.Debug().Raw("? UNION ? ORDER BY id ASC", previousQuery, nextQuery)
		finalQuery = db.gormDB.Raw(`
		(SELECT * FROM "messages" WHERE chat_id = ? AND id < ? ORDER BY id DESC LIMIT ?)
		UNION
		(SELECT * FROM "messages" WHERE chat_id = ? AND id > ? ORDER BY id ASC LIMIT ?)
		ORDER BY id ASC
		`,
			chatID, fromMsgID, count, // For previous
			chatID, fromMsgID, count, // For next
		)
	default:
		return nil, fmt.Errorf("invalid mode: %s", direction)
	}

	// Apply preloads and execute the query
	err := finalQuery.
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

	// For "both" or "previous", make sure results are ordered correctly
	// if direction == "previous" {
	// 	result = reverseMessages(result)
	// }

	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})
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
