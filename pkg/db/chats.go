package database

import (
	"api-channel/pkg/models"
	"api-channel/proto"
)

func GetChats(userId uint) (*proto.ChatsResponse, error) {

	var user models.User
	res := db.First(&user, userId).
		Preload("Chats").
		Preload("Replies").
		Preload("Replies.ReplyMessage").
		Preload("Replies.Thumbnails").
		Preload("Replies.Thumbnails.Thumbnail").
		Preload("Attachments").
		Preload("Attachments.Thumbnails", "type = ?", "placeholder").
		Find(&chats)

	if res.Error != nil {
		return nil, res.Error
	}

	var data []*proto.MessagesResponse_MessageData
	for _, m := range messages {

	}

	return &proto.MessagesResponse{Data: data}, nil

}
