package server

import (
	"api-channel/proto"
	"context"
)

func (s *Server) SendMessage(ctx context.Context, req *proto.SendMessageRequest) (*proto.SendMessageResponse, error) {
	token := req.Token

	// Check token is exist
	_, err := checkToken(token)
	if err != nil {
		return nil, err
	}

	return &proto.SendMessageResponse{}, nil
}
