package server

import (
	"api-channel/proto"
	"context"
)

func (s *Server) JoinGroup(ctx context.Context, req *proto.JoinGroupRequest) (*proto.JoinGroupResponse, error) {
	// Check token is exist
	t, err := checkToken(req.Token)
	if err != nil {
		return nil, err
	}

	err = s.db.CreateChatMember(uint(req.ChatId), uint32(t.UserID), "member")
	if err != nil {
		return nil, err
	}

	return &proto.JoinGroupResponse{}, nil
}
