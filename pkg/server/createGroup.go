package server

import (
	"api-channel/proto"
	"context"
)

func (s *Server) CreateGroup(ctx context.Context, req *proto.CreateGroupRequest) (*proto.CreateGroupResponse, error) {
	// Check token is exist
	t, err := checkToken(req.Token)
	if err != nil {
		return nil, err
	}

	chatId, err := s.db.CreateChat(true)
	if err != nil {
		return nil, err
	}

	groupId, err := s.db.CreateGroup(req.Name, uint(chatId), t.UserID)
	if err != nil {
		return nil, err
	}

	// Join current user to created group chat
	err = s.db.CreateChatMember(uint(chatId), uint32(t.UserID), "admin")
	if err != nil {
		return nil, err
	}

	return &proto.CreateGroupResponse{
		GroupId: groupId,
		ChatId:  chatId,
	}, nil
}
