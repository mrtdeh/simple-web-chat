package database

import (
	"api-channel/pkg/models"
	"api-channel/proto"
	"fmt"
	"sort"

	"gorm.io/gorm"
)

func isLastPage(id, pageSize, total int) bool {
	lastPage := (total-1)/pageSize + 1
	currentPage := (id-1)/pageSize + 1
	return currentPage == lastPage
}

func (db *ChatDatabase) GetLastMessageID(chatId uint32) (uint32, error) {
	var lastID uint32
	err := db.gormDB.Model(&models.Message{}).
		Select("id").
		Where("chat_id = ?", chatId).
		Order("id DESC").
		Limit(1).
		Scan(&lastID).Error
	return lastID, err
}

func (db *ChatDatabase) GetMessages(chatID, fromMsgID uint32, direction proto.GetMessagesRequest_Direction, count int32) ([]models.Message, error) {
	var result []models.Message
	var finalQuery *gorm.DB

	switch direction {
	case proto.GetMessagesRequest_PrevPage:
		// Only previous query
		finalQuery = db.gormDB.Table("messages").Select("*").
			Where("chat_id = ? AND id < ?", chatID, fromMsgID).
			Order("id DESC").
			Limit(int(count))

	case proto.GetMessagesRequest_NextPage:
		// Only next query
		finalQuery = db.gormDB.Table("messages").Select("*").
			Where("chat_id = ? AND id > ?", chatID, fromMsgID).
			Order("id ASC").
			Limit(int(count))

	case proto.GetMessagesRequest_LastPage:
		// Last messages
		finalQuery = db.gormDB.Debug().Table("messages").Select("*").
			Where("chat_id = ? AND id > ?", chatID, fromMsgID).
			Order("id DESC").
			Limit(int(count))

	case proto.GetMessagesRequest_BothPage:
		// Combine both with UNION
		// finalQuery = db.gormDB.Raw(fmt.Sprintf("(%s) UNION (%s)", previousQuery, nextQuery))
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
		return nil, fmt.Errorf("invalid mode: %v", direction)
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
