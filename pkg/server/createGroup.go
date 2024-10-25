package server

import (
	database "api-channel/pkg/db"
	"api-channel/pkg/models"
	"api-channel/proto"
	"context"
)

func (s *Server) CreateGroup(ctx context.Context, req *proto.CreateGroupRequest) (*proto.CreateGroupResponse, error) {
	// Check token is exist
	t, err := checkToken(req.Token)
	if err != nil {
		return nil, err
	}

	db := database.GetInstance()

	// Create chat for group
	chat := &models.Chat{
		IsGroup: true,
	}
	res := db.Create(chat)
	if res.Error != nil {
		return nil, res.Error
	}

	// Create group chat
	g := &models.Group{
		Name:    req.Name,
		OwnerID: t.UserID,
		ChatID:  chat.ID,
	}
	res = db.Create(g)
	if res.Error != nil {
		return nil, res.Error
	}

	// Join current user to created group chat
	_, err = s.JoinGroup(ctx, &proto.JoinGroupRequest{
		Token:  req.Token,
		ChatId: uint32(chat.ID),
		Role:   "admin",
	})
	if err != nil {
		return nil, err
	}

	return &proto.CreateGroupResponse{
		GroupId: uint32(g.ID),
		ChatId:  uint32(chat.ID),
	}, nil
}
