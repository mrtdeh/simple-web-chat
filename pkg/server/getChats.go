package server

import (
	"api-channel/proto"
	"context"
)

func (s *Server) GetChats(ctx context.Context, req *proto.GetChatsRequest) (*proto.ChatsResponse, error) {

	t, err := checkToken(req.Token)
	if err != nil {
		return nil, err
	}

	rows, err := s.db.GetChats(t.UserID)
	if err != nil {
		return nil, err
	}

	var Data []*proto.ChatsResponse_ChatData
	if err := s.db.ScanRows(rows, &Data); err != nil {
		return nil, err
	}

	return &proto.ChatsResponse{Data: Data}, nil
}
