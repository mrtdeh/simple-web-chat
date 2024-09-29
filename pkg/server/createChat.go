package server

import (
	database "api-channel/pkg/db"
	"api-channel/pkg/models"
	"api-channel/proto"
	"context"
)

func (s *Server) CreateChat(ctx context.Context, req *proto.CreateChatRequest) (*proto.CreateChatResponse, error) {

	db := database.GetInstance()

	// Insert message into messages table
	chat := models.Chat{
		Name:    req.Name,
		IsGroup: req.IsGroup,
	}
	db.Create(&chat)

	return &proto.CreateChatResponse{ChatId: uint32(chat.ID)}, nil
}
