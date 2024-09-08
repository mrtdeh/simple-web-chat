package server

import (
	"api-channel/proto"
	"context"
)

func (s *GRPCServer) Ping(ctx context.Context, req *proto.PingRequest) (*proto.PingResponse, error) {

	return &proto.PingResponse{}, nil
}
