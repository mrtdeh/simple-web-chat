package server

import (
	"api-channel/proto"
	"context"
	"fmt"
	"time"
)

func (s *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	user := req.Username
	pass := req.Password

	// TODO: check username and password in db
	if user != "admin" && pass != "admin" {
		return nil, fmt.Errorf("username or password incorrect")
	}

	token := tokenGenerator()
	tokens[token] = TokenData{
		Username:   user,
		ExpireTime: time.Now().Add(time.Minute * 5),
	}

	return &proto.LoginResponse{
		Token: token,
	}, nil
}
