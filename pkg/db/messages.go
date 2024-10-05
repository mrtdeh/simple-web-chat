package database

import (
	"api-channel/pkg/models"
	"api-channel/proto"
)

func GetMessagesWithReplies(chatID uint) (*proto.MessagesResponse, error) {

	var messages []models.Message
	res := db.Model(&models.Message{ChatID: 1}).
		Preload("Sender").
		Preload("Replies").
		Preload("Replies.ReplyMessage").
		Preload("Replies.Thumbnails").
		Preload("Replies.Thumbnails.Thumbnail").
		Preload("Attachments").
		Preload("Attachments.Thumbnails", "type = ?", "placeholder").
		Find(&messages)

	if res.Error != nil {
		return nil, res.Error
	}

	var data []*proto.MessagesResponse_MessageData
	for _, m := range messages {

		// Fetch message attachment's placeholder
		var attachsPlaceholder []string
		for _, att := range m.Attachments {
			for _, t := range att.Thumbnails {
				if t.Type == "placeholder" {
					attachsPlaceholder = append(attachsPlaceholder, t.Base64)
					break
				}
			}
		}

		// Fetch replied messages with thumbnails
		var repliedMessages []*proto.MessagesResponse_RepliedMessage
		for _, r := range m.Replies {
			var thumbnails []string
			for _, t := range r.Thumbnails {
				thumbnails = append(thumbnails, t.Thumbnail.Base64)
			}
			repliedMessages = append(repliedMessages, &proto.MessagesResponse_RepliedMessage{
				MessageId:  uint32(r.ReplyMessageId),
				Content:    r.ReplyMessage.Content,
				Thumbnails: thumbnails,
			})
		}

		// Collect message data
		data = append(data, &proto.MessagesResponse_MessageData{
			SenderId:        uint32(m.Sender.ID),
			MessageId:       uint32(m.ID),
			Content:         m.Content,
			SendAt:          m.CreatedAt.Unix(),
			Attachements:    attachsPlaceholder,
			RepliedMessages: repliedMessages,
		})

	}

	return &proto.MessagesResponse{Data: data}, nil

}
