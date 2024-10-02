package server

import (
	database "api-channel/pkg/db"
	"api-channel/pkg/models"
	"api-channel/proto"
	"context"
	"fmt"
)

func (s *Server) GetMessages(ctx context.Context, req *proto.GetMessagesRequest) (*proto.Messages, error) {

	db := database.GetInstance()

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
			fmt.Println("reply to :", r.ReplyMessageId)
			data[i].Replies = append(data[i].Replies, &proto.ReplayTo{
				Id: uint32(r.ReplyMessageId),
				// Type: r.Type,
			})
		}
	}

	return &proto.Messages{Data: data}, nil
}
