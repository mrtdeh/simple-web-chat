package server

import (
	database "api-channel/pkg/db"
	"api-channel/pkg/models"
	"api-channel/proto"
	"context"
	"time"
)

func (s *Server) AddMember(ctx context.Context, req *proto.AddMemberRequest) (*proto.AddMemberResponse, error) {
	// Check token is exist
	// t, err := checkToken(req.Token)
	// if err != nil {
	// 	return nil, err
	// }

	db := database.GetInstance()

	for _, uid := range req.MembersId {
		member := &models.ChatMember{
			UserID:   uint(uid),
			ChatID:   uint(req.ChatId),
			Role:     "member",
			JoinedAt: time.Now(),
		}
		res := db.Create(member)
		if res.Error != nil {
			return nil, res.Error
		}
	}

	return &proto.AddMemberResponse{Status: "ok"}, nil
}
