package server

import (
	"api-channel/proto"
	"context"
)

func (s *Server) GetMessages(ctx context.Context, req *proto.GetMessagesRequest) (*proto.MessagesResponse, error) {

	// db := database.GetInstance()

	// var messages []models.Message
	// res := db.Model(&models.Message{ChatID: 1}).Preload("Sender").Preload("Replies").Find(&messages)
	// if res.Error != nil {
	// 	return nil, res.Error
	// }

	// var data []*proto.MessageData
	// for i, m := range messages {
	// 	data = append(data, &proto.MessageData{
	// 		SenderId:  uint32(m.Sender.ID),
	// 		MessageId: uint32(m.ID),
	// 		Content:   m.Content,
	// 	})
	// 	for _, r := range m.Replies {
	// 		data[i].RepliedMessages = append(data[i].RepliedMessages, &proto.RepliedMessage{
	// 			MessageId: uint32(r.ReplyMessageId),
	// 		})
	// 	}
	// }

	// return &proto.Messages{Data: data}, nil
	return nil, nil
}
