package server

import (
	"api-channel/proto"
	"context"
	"fmt"
)

func (s *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {

	userId, err := s.db.CheckUser(req.Username, req.Password)
	if err != nil {
		return nil, fmt.Errorf("username or password incorrect")
	}

	// Create new token
	token := tm.Create(userId, req.Username)

	return &proto.LoginResponse{
		Token:  token,
		UserId: userId,
	}, nil
}
