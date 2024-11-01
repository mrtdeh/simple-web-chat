package server

import (
	"api-channel/proto"
	"context"
	"fmt"
)

func (s *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	// user := req.Username
	// pass := req.Password

	// db := database.GetInstance()
	// TODO: check username and password in db
	userId, err := s.db.CheckUser(req.Username, req.Password)
	if err != nil {
		return nil, fmt.Errorf("username or password incorrect")
	}

	// Create new token
	token := tm.Create(userId, req.Username)

	return &proto.LoginResponse{
		Token: token,
	}, nil
}
