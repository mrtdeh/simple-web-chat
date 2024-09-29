package server

import (
	database "api-channel/pkg/db"
	"api-channel/pkg/models"
	"api-channel/proto"
	"context"
)

func (s *Server) SendMessage(ctx context.Context, req *proto.MessageRequest) (*proto.MessageResponse, error) {
	token := req.Token

	// Check token is exist
	t, err := checkToken(token)
	if err != nil {
		return nil, err
	}

	db := database.GetInstance()

	// Insert message into messages table
	msg := models.Message{
		SenderID: t.UserID,
		ChatID:   uint(req.ChatId),
		Content:  req.Content,
	}
	db.Create(&msg)

	for _, r := range req.Replaies {
		db.Create(&models.Reply{
			MessageID: msg.ID,
			ReplyToID: uint(r.Id),
			Type:      r.Type,
		})
	}

	return &proto.MessageResponse{MessageId: uint32(msg.ID)}, nil
}
