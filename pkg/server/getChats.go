package server

import (
	database "api-channel/pkg/db"
	"api-channel/proto"
	"context"
)

func (s *Server) GetChats(ctx context.Context, req *proto.GetChatsRequest) (*proto.ChatsResponse, error) {
	t, err := checkToken(req.Token)
	if err != nil {
		return nil, err
	}

	return database.GetChats(t.UserID)
}
