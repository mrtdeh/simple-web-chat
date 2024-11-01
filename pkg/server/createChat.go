package server

import (
	"api-channel/proto"
	"context"
)

func (s *Server) CreateChat(ctx context.Context, req *proto.CreateChatRequest) (*proto.CreateChatResponse, error) {

	chatId, err := s.db.CreateChat(req.IsGroup)
	if err != nil {
		return nil, err
	}
	return &proto.CreateChatResponse{ChatId: chatId}, nil
}
