package server

import (
	database "api-channel/pkg/db"
	"api-channel/proto"
	"context"
)

func (s *Server) GetChats(ctx context.Context, req *proto.GetChatsRequest) (*proto.ChatsResponse, error) {

	return database.GetChats(uint(req.UserId))
}
