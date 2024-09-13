package server

import (
	"api-channel/proto"
	"context"
)

func (s *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {

	return &proto.LoginResponse{}, nil
}
