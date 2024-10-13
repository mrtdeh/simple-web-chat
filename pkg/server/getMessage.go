package server

import (
	database "api-channel/pkg/db"
	"api-channel/proto"
	"context"
)

func (s *Server) GetMessages(ctx context.Context, req *proto.GetMessagesRequest) (*proto.MessagesResponse, error) {

	// return &proto.Messages{Data: data}, nil
	return database.GetMessagesWithReplies(uint(req.ChatId))
}
