package database

import (
	"api-channel/pkg/models"
	"api-channel/proto"
	"fmt"
)

type Message struct {
	MessageId   uint   `json:"message_id"`
	Content     string `json:"content"`
	SenderId    uint   `json:"sender_id"`
	ChatId      uint   `json:"chat_id"`
	SendDate    string `json:"send_date"`
	Attachments []Attachment
	Replies     []MessageReply
}
type MessageReply struct {
	Message    *Message
	Attachment *Attachment
}
type Attachment struct {
	AttachmentId uint   `json:"attachment_id"`
	Data         string `json:"data"`
}

func GetMessagesWithReplies(chatID uint) ([]Message, error) {

	var messages []models.Message
	res := db.Model(&models.Message{ChatID: 1}).Preload("Sender").Preload("Replies").Find(&messages)
	if res.Error != nil {
		return nil, res.Error
	}

	var data []*proto.MessageData
	for i, m := range messages {
		data = append(data, &proto.MessageData{
			SenderId:  uint32(m.Sender.ID),
			MessageId: uint32(m.ID),
			Content:   m.Content,
		})
		fmt.Println("msg id : ", m.ID)
		for _, r := range m.Replies {
			fmt.Println("reply to :", r.ReplyToID)
			data[i].Replies = append(data[i].Replies, &proto.ReplayTo{
				Id:   uint32(r.ReplyToID),
				Type: r.Type,
			})
		}
	}

}

// func GetMessagesWithReplies(chatID uint) {
// 	var messages []models.Message
// 	db.Where("chat_id = ?", chatID).Find(&messages)

// 	for _, message := range messages {
// 		fmt.Printf("Message: %s", message.Content)

// 		// دریافت پاسخ‌های مرتبط با پیام‌های متنی
// 		var textReplies []models.Reply
// 		db.Where("message_id = ? AND type = 'text'", message.ID).Find(&textReplies)
// 		for _, reply := range textReplies {
// 			var replyMessage models.Message
// 			db.First(&replyMessage, reply.ReplyToID)
// 			fmt.Printf(" (replied to text: %s)", replyMessage.Content)
// 		}

// 		// دریافت پاسخ‌های مرتبط با فایل‌ها
// 		var fileReplies []models.Reply
// 		db.Where("message_id = ? AND type = 'file'", message.ID).Find(&fileReplies)
// 		for _, reply := range fileReplies {
// 			var attachment models.Attachment
// 			db.First(&attachment, reply.ReplyToID)
// 			fmt.Printf(" (replied to file: %s)", attachment.FilePath)
// 		}
// 	}
// }
