package database

import (
	"api-channel/pkg/models"
	"fmt"
)

func GetMessagesWithReplies(chatID uint) {
	var messages []models.Message
	db.Where("chat_id = ?", chatID).Find(&messages)

	for _, message := range messages {
		fmt.Printf("Message: %s", message.Content)

		// دریافت پاسخ‌های مرتبط با پیام‌های متنی
		var textReplies []models.Reply
		db.Where("message_id = ? AND type = 'text'", message.ID).Find(&textReplies)
		for _, reply := range textReplies {
			var replyMessage models.Message
			db.First(&replyMessage, reply.ReplyToID)
			fmt.Printf(" (replied to text: %s)", replyMessage.Content)
		}

		// دریافت پاسخ‌های مرتبط با فایل‌ها
		var fileReplies []models.Reply
		db.Where("message_id = ? AND type = 'file'", message.ID).Find(&fileReplies)
		for _, reply := range fileReplies {
			var attachment models.Attachment
			db.First(&attachment, reply.ReplyToID)
			fmt.Printf(" (replied to file: %s)", attachment.FilePath)
		}
	}
}
