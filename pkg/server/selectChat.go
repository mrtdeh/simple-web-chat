package server

import (
	"api-channel/proto"
	"context"
)

func (s *Server) SelectChat(ctx context.Context, req *proto.SelectChatRequest) (*proto.EmptyRequest, error) {
	token := req.Token

	// Check token is exist
	t, err := checkToken(token)
	if err != nil {
		return nil, err
	}

	session := s.Sessions.Get(t.Username)
	session.activeChat(req.ChatId)

	return &proto.EmptyRequest{}, nil
}
