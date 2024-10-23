package server

import (
	database "api-channel/pkg/db"
	"api-channel/pkg/models"
	"api-channel/proto"
	"context"
	"time"
)

func (s *Server) JoinGroup(ctx context.Context, req *proto.JoinGroupRequest) (*proto.JoinGroupResponse, error) {
	// Check token is exist
	t, err := checkToken(req.Token)
	if err != nil {
		return nil, err
	}

	db := database.GetInstance()

	member := &models.ChatMember{
		UserID:   t.UserID,
		ChatID:   uint(req.ChatId),
		Role:     "member",
		JoinedAt: time.Now(),
	}
	res := db.Create(member)
	if res.Error != nil {
		return nil, res.Error
	}

	return &proto.JoinGroupResponse{}, nil
}
