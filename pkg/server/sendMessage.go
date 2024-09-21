package server

import (
	database "api-channel/pkg/db"
	"api-channel/pkg/models"
	"api-channel/proto"
	"context"
	"time"
)

func (s *Server) SendMessage(ctx context.Context, req *proto.MessageRequest) (*proto.MessageResponse, error) {
	token := req.Token

	// Check token is exist
	_, err := checkToken(token)
	if err != nil {
		return nil, err
	}

	db := database.GetInstance()
	// Check for incomming text message

	// Insert message into messages table
	db.Create(models.Message{
		// SenderID: fetch sender id from token store,
		ChatID:    uint(req.ChatId),
		Content:   req.Content,
		CreatedAt: time.Now(),
	})

	return &proto.MessageResponse{}, nil
}
