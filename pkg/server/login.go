package server

import (
	"api-channel/proto"
	"context"
	"fmt"
)

func (s *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	user := req.Username
	pass := req.Password

	// TODO: check username and password in db
	if user != "admin" && pass != "admin" {
		return nil, fmt.Errorf("username or password incorrect")
	}

	// Create new token
	token := tm.Create(user)

	return &proto.LoginResponse{
		Token: token,
	}, nil
}
